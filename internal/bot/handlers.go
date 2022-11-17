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

type NewCTFPayload struct {
	Name      string  `optname:"name"`
	URL       string  `optname:"url"`
	Password  *string `optname:"password"`
	Username  *string `optname:"username"`
	CTFTimeID *int    `optname:"ctftime-id"`
}

type NewChallPayload struct {
	Name string `optname:"name"`
}

type SolvePayload struct {
	Flag string `optname:"flag"`
}

func (b *bot) newCTF(m *discordgo.InteractionCreate, p *NewCTFPayload) {
	b.log.Info("new ctf", zap.Any("p", p))

	exists, err := models.CTFChannels(
		models.CTFChannelWhere.Title.EQ(p.Name),
	).Exists(context.TODO(), b.db)
	if err != nil {
		b.log.Error("db err", zap.Error(err))
		return
	}

	if exists {
		// TODO ask if we should upsert
		b.reply(m.Interaction, fmt.Sprintf("`%s` already exists", p.Name))
		return
	}

	category, err := b.sess.GuildChannelCreate(m.GuildID, p.Name, discordgo.ChannelTypeGuildCategory)
	if err != nil {
		b.log.Error("could not create category", zap.Error(err))
		return
	}

	ctf := &models.CTFChannel{
		ID:        category.ID,
		GuildID:   m.GuildID,
		Title:     p.Name,
		URL:       p.URL,
		Username:  null.StringFromPtr(p.Username),
		CtftimeID: null.IntFromPtr(p.CTFTimeID),
		Password:  null.StringFromPtr(p.Password),
	}

	chann, err := b.sess.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:     p.Name,
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

	err = b.sess.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Created %s", chann.Mention()),
		},
	})
	if err != nil {
		b.log.Error("could not respond", zap.Error(err))
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

	err = b.sess.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "done!",
		},
	})

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

	err = b.sess.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("nice, flag: `%s`", p.Flag),
		},
	})
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
