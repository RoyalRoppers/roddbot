package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/models"
	"github.com/movitz-s/roddbot/internal/permissions"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type NewUpdateEventPayload struct {
	Title               *string `optname:"title"`
	Start               *string `optname:"start"`
	End                 *string `optname:"end"`
	Location            *string `optname:"location"`
	Description         *string `optname:"description"`
	AnnouncementChannel *string `optname:"announcement-channel"`
	AnnouncementTime    *string `optname:"announcement-time"`
}

const DateTimeFormat = "2006-01-02 15:04"

func eventEmbed(e models.Event) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       e.Title,
		Description: e.Description,
		Color:       0xf9c867,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "When?",
				Value: fmt.Sprintf(
					"<t:%d:f> - <t:%d:f>",
					e.StartTime.Unix(),
					e.EndTime.Unix(),
				),
				Inline: true,
			},
			{
				Name:   "Where?",
				Value:  e.Location,
				Inline: true,
			},
		},
	}
}

func (b *bot) newEvent(m *discordgo.InteractionCreate, p *NewUpdateEventPayload) {
	if !b.hasPermission(m, permissions.EventCreate) {
		b.reply(m.Interaction, "No permission to create Event")
		return
	}

	event := &models.Event{
		Title:               *p.Title,
		Location:            *p.Location,
		Description:         *p.Description,
		AnnouncementChannel: *p.AnnouncementChannel,
	}

	var err error
	event.StartTime, err = time.Parse(DateTimeFormat, *p.Start)
	if err != nil {
		b.log.Info("invalid start time format", zap.Error(err),
			zap.Any("start", p.Start))
		b.reply(m.Interaction,
			"Invalid start time format! Want format: 2006-01-02 15:04")
		return
	}

	event.EndTime, err = time.Parse(DateTimeFormat, *p.End)
	if err != nil {
		b.log.Info("invalid end time format", zap.Error(err),
			zap.Any("end", p.End))
		b.reply(m.Interaction,
			"Invalid end time format! Want format: 2006-01-02 15:04")
		return
	}

	now := time.Now()
	if now.After(event.StartTime) {
		b.log.Info("start time must be after now", zap.Any("now", now),
			zap.Any("start", event.StartTime))
		b.reply(m.Interaction, "Start time must be after now")
		return
	}

	if event.StartTime.After(event.EndTime) {
		b.log.Info("start time must be before end time",
			zap.Any("start", event.StartTime), zap.Any("end", event.EndTime))
		b.reply(m.Interaction, "Start time must be before end time")
		return
	}

	if p.AnnouncementTime != nil {
		event.AnnouncementTime.Valid = true
		event.AnnouncementTime.Time, err = time.Parse(DateTimeFormat,
			*p.AnnouncementTime)
		if err != nil {
			b.log.Info("invalid announcement time format", zap.Error(err),
				zap.Any("announcement-time", p.AnnouncementTime))
			b.reply(m.Interaction,
				"Invalid announcement time format! Want format: 2006-01-02 15:04")
			return
		}

		if now.After(event.AnnouncementTime.Time) {
			b.log.Info("announcement time must be after now",
				zap.Any("now", now),
				zap.Any("announcement-time", event.AnnouncementTime.Time))
			b.reply(m.Interaction, "Announcement time time must be after now")
			return
		}

		if event.AnnouncementTime.Time.After(event.StartTime) {
			b.log.Info("announcement time must be before start time",
				zap.Any("announcement-time", event.AnnouncementTime.Time),
				zap.Any("start", event.StartTime))
			b.reply(m.Interaction,
				"Announcement time time must be before start time")
			return
		}
	} else {
		msg, err := b.sess.ChannelMessageSendComplex(event.AnnouncementChannel,
			&discordgo.MessageSend{
				Content: "@everyone",
				Embeds:  []*discordgo.MessageEmbed{eventEmbed(*event)},
			})
		if err != nil {
			b.log.Error("could not send announcement", zap.Error(err))
			b.reply(m.Interaction, "Could not send announcement")
			return
		}
		event.AnnouncementMSGID.SetValid(msg.ID)

		e, err := b.sess.GuildScheduledEventCreate(m.GuildID,
			&discordgo.GuildScheduledEventParams{
				Name:               event.Title,
				Description:        event.Description,
				ScheduledStartTime: &event.StartTime,
				ScheduledEndTime:   &event.EndTime,
				PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
				EntityType:         3,
				EntityMetadata:     &discordgo.GuildScheduledEventEntityMetadata{Location: event.Location},
			},
		)
		if err != nil {
			b.log.Error("could not create event", zap.Error(err))
			b.reply(m.Interaction, "Could not create event")
			return
		}
		event.AnnouncementEventID.SetValid(e.ID)
	}

	err = event.Insert(context.TODO(), b.db, boil.Infer())
	if err != nil {
		b.log.Error("could not insert event", zap.Error(err), zap.Any("event", event))
		b.reply(m.Interaction, "DB failed => state mismatch, chaos")
		return
	}

	err = b.reply(m.Interaction, fmt.Sprintf("Created event %s", event.Title))
	if err != nil {
		b.log.Error("could not respond", zap.Error(err))
	}
}

func (b *bot) listEvents(m *discordgo.InteractionCreate) {
	if !b.hasPermission(m, permissions.EventCreate) {
		b.reply(m.Interaction, "No permission to list events")
		return
	}

	events, err := models.Events().All(context.TODO(), b.db)
	if err != nil {
		b.log.Error("could not fetch events", zap.Error(err))
		b.reply(m.Interaction, "DB failed => chaos")
		return
	}
	// TODO: continue with adding an update function.
	var embeds []*discordgo.MessageEmbed
	for _, event := range events {
		e := eventEmbed(*event)
		e.Author = &discordgo.MessageEmbedAuthor{
			Name: fmt.Sprintf("Event id: %d", event.ID),
		}
		embeds = append(embeds, e)
	}
	b.sess.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: "Here are the events:",
		Embeds:  embeds,
	})
	if err != nil {
		b.log.Error("could not respond", zap.Error(err))
	}
}
