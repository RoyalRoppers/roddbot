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

func channelPosition(category *int, solved bool) int {

	/*
		reserved positions

		1: ctf channel
		10: challs without category
		11-100: challs with category
		1000: solved challs without category
		1001-1100: solved challs with category
		2000: voice chat
	*/

	pos := 10

	if solved {
		pos = 1000
	}

	if category != nil {
		pos += *category + 1
	}

	return pos
}

func channelName(title string, category *int, solved bool) string {
	name := title
	if category != nil {
		name = translateCategory(*category) + "-" + name
	}
	if solved {
		name = "✔" + name
	}
	return name
}

func translateCategory(x int) string {
	switch x {
	case 0:
		return "web"
	case 1:
		return "rev"
	case 2:
		return "pwn"
	case 3:
		return "crypto"
	case 4:
		return "forensics"
	case 5:
		return "misc"
	case 6:
		return "blockchain"
	case 7:
		return "osint"
	default:
		return ""
	}
}
