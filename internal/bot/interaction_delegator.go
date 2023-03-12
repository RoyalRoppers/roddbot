package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/unmarshal"
	"go.uber.org/zap"
)

func (b *bot) msgCreateHandler(s *discordgo.Session, m *discordgo.InteractionCreate) {
	defer func() {
		err := recover()
		if err != nil {
			b.log.Error("interaction handler paniced", zap.Any("err", err))

			go func() {
				errx := b.reply(m.Interaction, `¯\_( ͡🔥 ͜ʖ ͡🔥)_/¯`)
				if errx != nil {
					b.log.Error("panic response failed lol", zap.Error(errx))
				}
			}()
		}
	}()

	if !b.TryLock() {
		b.reply(m.Interaction, "Retry when other commands are done running.")
		return
	}
	defer b.Unlock()

	d := m.Interaction.ApplicationCommandData()

	_, err := b.guildSanityCheck(m)
	if err != nil {
		return
	}

	b.log.Info("command received", zap.Any("payload", d))

	switch d.Name {
	case "map-roles":
		var payload RoleMapPayload
		unmarshal.Unmarshal(d.Options, &payload)
		b.handlePermissionMap(m, &payload)

	case "ctf":
		b.handleCTF(m, d)

	case "events":
		b.handleEvents(m, d)

	default:
		b.log.Warn("unknown command", zap.String("cmd", d.Name))
	}
}

func (b *bot) handleCTF(m *discordgo.InteractionCreate, d discordgo.ApplicationCommandInteractionData) {
	switch d.Options[0].Name {
	case "new":
		var payload NewUpdateCTFPayload
		unmarshal.Unmarshal(d.Options[0].Options, &payload)
		b.newCTF(m, &payload)

	case "update":
		var payload NewUpdateCTFPayload
		unmarshal.Unmarshal(d.Options[0].Options, &payload)
		b.updateCTF(m, &payload)

	case "solve":
		var payload SolvePayload
		unmarshal.Unmarshal(d.Options[0].Options, &payload)
		b.solve(m, &payload)

	case "chall":
		var payload NewChallPayload
		unmarshal.Unmarshal(d.Options[0].Options, &payload)
		b.newChall(m, &payload)

	case "import-ctfd":
		b.importCtfd(m)

	case "purge":
		b.purge(m)

	default:
		b.log.Error("unhandeled interaction", zap.Any("interaction", d))
		b.reply(m.Interaction, "🚨🚨 Unkown interaction, something is wrong 🚨🚨")
	}
}

func (b *bot) handleEvents(m *discordgo.InteractionCreate, d discordgo.ApplicationCommandInteractionData) {
	switch d.Options[0].Name {
	case "new":
		var payload NewUpdateEventPayload
		unmarshal.Unmarshal(d.Options[0].Options, &payload)
		b.newEvent(m, &payload)

	case "list":
		b.listEvents(m)

	default:
		b.log.Error("unhandeled interaction", zap.Any("interaction", d))
		b.reply(m.Interaction, "🚨🚨 Unkown interaction, something is wrong 🚨🚨")
	}
}
