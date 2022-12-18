package bot

import "github.com/bwmarrin/discordgo"

func cmds() []*discordgo.ApplicationCommand {
	f := false
	return []*discordgo.ApplicationCommand{
		{
			Name:         "ctf",
			Type:         discordgo.ChatApplicationCommand,
			Description:  "ctf",
			DMPermission: &f,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "new",
					Description: "New CTF",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "name",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "The name of the CTF",
							Required:    true,
						},
						{
							Name:        "url",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "URL to the CTF",
							Required:    true,
						},
						{
							Name:        "username",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Username",
						},
						{
							Name:        "password",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Password",
						},
						{
							Name:        "ctftime-id",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Description: "CTFTime ID",
						},
						{
							Name:        "ctfd-api-token",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "CTFd API Token",
						},
					},
				},
				{
					Name:        "update",
					Description: "Update CTF",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "name",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "The name of the CTF",
						},
						{
							Name:        "url",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "URL to the CTF",
						},
						{
							Name:        "username",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Username",
						},
						{
							Name:        "password",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Password",
						},
						{
							Name:        "ctftime-id",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Description: "CTFTime ID",
						},
						{
							Name:        "ctfd-api-token",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "CTFd API Token",
						},
					},
				},
				{
					Name:        "chall",
					Description: "New CTF Challenge",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "name",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "The name of the challenge",
							Required:    true,
						},
					},
				},
				{
					Name:        "solve",
					Description: "Solve challenge",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "flag",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "The flag",
							Required:    true,
						},
					},
				},
			},
		},
	}
}
