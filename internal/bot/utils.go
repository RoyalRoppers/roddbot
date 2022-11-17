package bot

import "github.com/bwmarrin/discordgo"

func (b *bot) reply(i *discordgo.Interaction, msg string) error {
	return b.sess.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}
