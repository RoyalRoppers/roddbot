package bot

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func (b *bot) createChallChan(ctf *models.CTFChannel, guildID, name string) (*discordgo.Channel, error) {
	disChan, err := b.sess.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name:     name,
		Type:     discordgo.ChannelTypeGuildText,
		Topic:    channelTopic(ctf),
		ParentID: ctf.ID,
		Position: 25,
	})
	if err != nil {
		b.log.Error("could not create channel", zap.Error(err))
		return nil, err
	}

	challChan := &models.ChallChannel{
		ID:       disChan.ID,
		ParentID: ctf.ID,
		Title:    name,
	}

	err = challChan.Insert(context.TODO(), b.db, boil.Infer())
	if err != nil {
		b.log.Error("could not insert chall chan", zap.Error(err))
		return nil, err
	}

	return disChan, nil
}
