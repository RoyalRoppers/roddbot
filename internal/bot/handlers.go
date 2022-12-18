package bot

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
		// TODO ask if we should upsert
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

	disChan, err := b.sess.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:     p.Name,
		Type:     discordgo.ChannelTypeGuildText,
		Topic:    channelTopic(ctf),
		ParentID: ctf.ID,
		Position: 25,
	})
	if err != nil {
		b.log.Error("could not create channel", zap.Error(err))
		return
	}

	challChan := &models.ChallChannel{
		ID:       disChan.ID,
		ParentID: ctf.ID,
		Title:    p.Name,
	}

	err = challChan.Insert(context.TODO(), b.db, boil.Infer())
	if err != nil {
		b.log.Error("could not insert chall chan", zap.Error(err))
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
