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

	d := m.Interaction.ApplicationCommandData()

	if d.Name != "ctf" {
		return
	}

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

	default:
		b.log.Error("unhandeled interaction", zap.Any("interaction", d))
		b.reply(m.Interaction, "🚨🚨 Unkown interaction, something is wrong 🚨🚨")
	}
}
