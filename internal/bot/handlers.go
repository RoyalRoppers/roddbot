package bot

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/ctfd"
	"github.com/movitz-s/roddbot/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type NewUpdateCTFPayload struct {
	Name         *string `optname:"name"`
	URL          *string `optname:"url"`
	Password     *string `optname:"password"`
	Username     *string `optname:"username"`
	CTFTimeID    *int    `optname:"ctftime-id"`
	CTFDAPIToken *string `optname:"ctfd-api-token"`
}

type NewChallPayload struct {
	Name string `optname:"name"`
}

type SolvePayload struct {
	Flag string `optname:"flag"`
}

func (b *bot) newCTF(m *discordgo.InteractionCreate, p *NewUpdateCTFPayload) {
	b.log.Info("new ctf", zap.Any("p", p))

	exists, err := models.CTFChannels(
		models.CTFChannelWhere.Title.EQ(*p.Name),
	).Exists(context.TODO(), b.db)
	if err != nil {
		b.log.Error("db err", zap.Error(err))
		return
	}

	if exists {
		b.reply(m.Interaction, fmt.Sprintf("`%s` already exists, please update it instead", *p.Name))
		return
	}

	category, err := b.sess.GuildChannelCreate(m.GuildID, *p.Name, discordgo.ChannelTypeGuildCategory)
	if err != nil {
		b.log.Error("could not create category", zap.Error(err))
		return
	}

	ctf := &models.CTFChannel{
		ID:        category.ID,
		GuildID:   m.GuildID,
		Title:     *p.Name,
		URL:       *p.URL,
		Username:  null.StringFromPtr(p.Username),
		CtftimeID: null.IntFromPtr(p.CTFTimeID),
		Password:  null.StringFromPtr(p.Password),
		APIToken:  null.StringFromPtr(p.CTFDAPIToken),
	}

	chann, err := b.sess.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:     *p.Name,
		Type:     discordgo.ChannelTypeGuildText,
		Topic:    channelTopic(ctf),
		ParentID: category.ID,
		Position: 1,
	})
	if err != nil {
		b.log.Error("could not create channel", zap.Error(err))
		return
	}

	ctf.TopicChan = chann.ID
	err = ctf.Insert(context.TODO(), b.db, boil.Infer())
	if err != nil {
		b.log.Error("could not insert ctf", zap.Error(err), zap.Any("ctf", ctf))
		b.reply(m.Interaction, "DB failed => state mismatch, chaos")
		return
	}

	err = b.reply(m.Interaction, fmt.Sprintf("Created %s", chann.Mention()))
	if err != nil {
		b.log.Error("could not respond", zap.Error(err))
	}
}

func (b *bot) updateCTF(m *discordgo.InteractionCreate, p *NewUpdateCTFPayload) {
	b.log.Info("update ctf", zap.Any("p", p))

	ctf, err := models.CTFChannels(
		models.CTFChannelWhere.TopicChan.EQ(m.ChannelID),
	).One(context.TODO(), b.db)
	if err == sql.ErrNoRows {
		b.reply(m.Interaction, "This is not a topic channel for a CTF, please run this command in the appropriate channel")
		return
	}
	if err != nil {
		b.log.Error("db err", zap.Error(err))
		return
	}

	cols := models.M{}

	if p.Name != nil {
		cols[models.CTFChannelColumns.Title] = *p.Name
	}
	if p.CTFTimeID != nil {
		cols[models.CTFChannelColumns.CtftimeID] = *p.CTFTimeID
	}
	if p.Password != nil {
		cols[models.CTFChannelColumns.Password] = *p.Password
	}
	if p.URL != nil {
		cols[models.CTFChannelColumns.URL] = *p.URL
	}
	if p.Username != nil {
		cols[models.CTFChannelColumns.Username] = *p.Username
	}
	if p.CTFDAPIToken != nil {
		cols[models.CTFChannelColumns.APIToken] = *p.CTFDAPIToken
	}

	if len(cols) == 0 {
		err = b.reply(m.Interaction, "Nothing to update")
		if err != nil {
			b.log.Error("could not reply", zap.Error(err))
		}
		return
	}

	_, err = models.CTFChannels(
		models.CTFChannelWhere.ID.EQ(ctf.ID),
	).UpdateAll(
		context.TODO(),
		b.db,
		cols,
	)
	if err != nil {
		b.log.Error("db err", zap.Error(err))
		return
	}

	ctf, err = models.CTFChannels(
		models.CTFChannelWhere.TopicChan.EQ(m.ChannelID),
	).One(context.TODO(), b.db)
	if err != nil {
		b.log.Error("db err", zap.Error(err))
		return
	}

	checkRate := func(err error) {
		if err, ok := err.(*discordgo.RateLimitError); ok {
			b.reply(m.Interaction, "Stop updating so often, you just hit a rate limit. Retry after "+err.RetryAfter.String()+"\n(There will be a mismatch in the descriptions now)")
		}
	}

	if p.Name != nil {
		_, err = b.sess.ChannelEditComplex(ctf.ID, &discordgo.ChannelEdit{
			Name: *p.Name,
		})
		if err != nil {
			checkRate(err)
			b.log.Error("could not update category", zap.Error(err))
			return
		}
		_, err = b.sess.ChannelEditComplex(ctf.TopicChan, &discordgo.ChannelEdit{
			Name: *p.Name,
		})
		if err != nil {
			checkRate(err)
			b.log.Error("could not update topic", zap.Error(err))
			return
		}
	}
	if p.CTFTimeID != nil || p.Password != nil || p.URL != nil || p.Username != nil {
		chans, err := models.ChallChannels(
			models.ChallChannelWhere.ParentID.EQ(ctf.ID),
		).All(context.TODO(), b.db)
		if err != nil {
			b.log.Error("could not list chall chans", zap.Error(err))
			return
		}

		_, err = b.sess.ChannelEditComplex(ctf.TopicChan, &discordgo.ChannelEdit{
			Topic: channelTopic(ctf),
		})
		if err != nil {
			checkRate(err)
			b.log.Error("could not update topic", zap.Error(err))
			return
		}

		for _, v := range chans {
			_, err = b.sess.ChannelEditComplex(v.ID, &discordgo.ChannelEdit{
				Topic: channelTopic(ctf),
			})
			if err != nil {
				checkRate(err)
				b.log.Error("could not update chall chan", zap.Error(err))
				return
			}
		}
	}

	err = b.reply(m.Interaction, "Updated!")
	if err != nil {
		b.log.Error("could not reply", zap.Error(err))
	}
}

func (b *bot) newChall(m *discordgo.InteractionCreate, p *NewChallPayload) {
	ctf, err := models.CTFChannels(
		models.CTFChannelWhere.TopicChan.EQ(m.ChannelID),
	).One(context.TODO(), b.db)
	if err == sql.ErrNoRows {
		b.reply(m.Interaction, "This is not a CTF channel, please act accordingly.")
		return
	}
	if err != nil {
		b.log.Error("could not check existance", zap.Error(err))
		b.reply(m.Interaction, "db err")
		return
	}

	_, err = b.createChallChan(ctf, m.GuildID, p.Name)
	if err != nil {
		return
	}

	err = b.reply(m.Interaction, "Done!")

	if err != nil {
		b.log.Error("could not respond", zap.Error(err))
	}
}

func (b *bot) solve(m *discordgo.InteractionCreate, p *SolvePayload) {
	challChan, err := models.ChallChannels(
		models.ChallChannelWhere.ID.EQ(m.ChannelID),
	).One(context.TODO(), b.db)
	if err == sql.ErrNoRows {
		b.reply(m.Interaction, "This is not a challenge channel, please act accordingly.")
		return
	}
	if err != nil {
		b.log.Error("could not get chall chan", zap.Error(err))
		b.reply(m.Interaction, "db err")
		return
	}

	_, err = b.sess.ChannelEdit(challChan.ID, &discordgo.ChannelEdit{
		Name:     "solved-" + challChan.Title,
		Position: 100,
	})
	if err != nil {
		b.log.Error("could not edit channel", zap.Error(err))
		return
	}

	challChan.SolvedAt = null.TimeFrom(time.Now())
	challChan.Flag = null.StringFrom(p.Flag)
	_, err = challChan.Update(context.TODO(), b.db, boil.Whitelist(models.ChallChannelColumns.SolvedAt, models.ChallChannelColumns.Flag))
	if err != nil {
		b.log.Error("could not update chall chan", zap.Error(err))
		return
	}

	err = b.reply(m.Interaction, fmt.Sprintf("nice, flag: `%s`", p.Flag))
	if err != nil {
		b.log.Error("could not respond", zap.Error(err))
	}
}

func (b *bot) importCtfd(m *discordgo.InteractionCreate) {
	ctf, err := models.CTFChannels(
		models.CTFChannelWhere.GuildID.EQ(m.GuildID),
	).One(context.TODO(), b.db)
	if err != nil {
		b.reply(m.Interaction, "Could not find guild")
		b.log.Error("could not get guild", zap.Error(err))
		return
	}

	if !ctf.APIToken.Valid {
		b.reply(m.Interaction, "Could not find a CTFd API token :(")
		return
	}

	c := ctfd.New(ctf.URL, ctf.APIToken.String)
	challs, err := c.GetChallanges()
	if err == ctfd.ErrBadAuth {
		b.reply(m.Interaction, "Could not authenticate with CTFd")
		return
	}
	if err != nil {
		b.reply(m.Interaction, "Something went wrong when talking to CTFd")
		return
	}

	err = b.sess.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		b.log.Error("could not respond", zap.Error(err))
		return
	}

	challIDs, err := models.ChallChannels(
		models.ChallChannelWhere.ParentID.EQ(ctf.ID),
		models.ChallChannelWhere.CTFDID.IsNotNull(),
		qm.Select(models.ChallChannelColumns.CTFDID),
	).All(context.TODO(), b.db)
	if err != nil {
		b.log.Error("could not get ctfd ids", zap.Error(err))
		return
	}

	msg := "Created channels:\n"
	for _, chall := range challs {
		for _, v := range challIDs {
			if v.CTFDID.Int == chall.ID {
				continue
			}
		}

		disChan, err := b.createChallChan(ctf, m.GuildID, chall.Name)
		if err != nil {
			return
		}
		msg += disChan.Mention() + "\n"
	}

	_, err = b.sess.InteractionResponseEdit(m.Interaction, &discordgo.WebhookEdit{
		Content: &msg,
	})
	if err != nil {
		b.log.Error("could not update response", zap.Error(err))
	}
}

func (b *bot) purge(m *discordgo.InteractionCreate) {
	ctf, err := models.CTFChannels(
		models.CTFChannelWhere.TopicChan.EQ(m.ChannelID),
		qm.Load(models.CTFChannelRels.ParentChallChannels),
	).One(context.TODO(), b.db)
	if err != nil {
		b.reply(m.Interaction, "Could not topic channel")
		b.log.Error("could not get guild", zap.Error(err))
		return
	}

	for _, v := range ctf.R.ParentChallChannels {
		_, err = b.sess.ChannelDelete(v.ID)
		if err != nil {
			b.log.Error("could not delete in discord", zap.Error(err), zap.String("challID", v.ID))
			return
		}
		_, err = v.Delete(context.TODO(), b.db) // im lazy and i dont care
		if err != nil {
			b.log.Error("could not delete chall chan", zap.Error(err), zap.String("challID", v.ID))
			return
		}
	}

	_, err = b.sess.ChannelDelete(ctf.TopicChan)
	if err != nil {
		b.log.Error("could not delete in discord", zap.Error(err), zap.String("challID", ctf.TopicChan))
		return
	}
	_, err = b.sess.ChannelDelete(ctf.ID)
	if err != nil {
		b.log.Error("could not delete in discord", zap.Error(err), zap.String("challID", ctf.ID))
		return
	}

	_, err = ctf.Delete(context.TODO(), b.db)
	if err != nil {
		b.log.Error("could not delete chall chan", zap.Error(err), zap.String("ctfID", ctf.ID))
		return
	}

	b.reply(m.Interaction, "deleted")
}

func channelTopic(p *models.CTFChannel) string {
	res := "url = " + p.URL

	if p.Username.Valid {
		res += "\nusername = " + p.Username.String
	}
	if p.Password.Valid {
		res += "\npassword = " + p.Password.String
	}
	if p.CtftimeID.Valid {
		res += fmt.Sprintf("\nctftime = https://ctftime.org/event/%d", p.CtftimeID.Int)
	}

	return res
}
