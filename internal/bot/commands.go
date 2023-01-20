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
						{
							Name:        "category",
							Type:        discordgo.ApplicationCommandOptionInteger,
							Description: "Category of the challenge. web/pwn/rev etc...",
							Choices: []*discordgo.ApplicationCommandOptionChoice{
								{
									Name:  "web",
									Value: 0,
								},
								{
									Name:  "rev",
									Value: 1,
								},
								{
									Name:  "pwn",
									Value: 2,
								},
								{
									Name:  "crypto",
									Value: 3,
								},
								{
									Name:  "forensics",
									Value: 4,
								},
								{
									Name:  "misc",
									Value: 5,
								},
								{
									Name:  "blockchain",
									Value: 6,
								},
								{
									Name:  "osint",
									Value: 7,
								},
							},
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
				{
					Name:        "import-ctfd",
					Description: "Import from CTFD",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options:     []*discordgo.ApplicationCommandOption{},
				},
				{
					Name:        "purge",
					Description: "Purges a group of CTF channels",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options:     []*discordgo.ApplicationCommandOption{},
				},
			},
		},
	}
}
