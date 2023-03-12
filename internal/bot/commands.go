package bot

import "github.com/bwmarrin/discordgo"

func cmds() []*discordgo.ApplicationCommand {
	f := false
	admin := int64(discordgo.PermissionAdministrator)
	return []*discordgo.ApplicationCommand{
		{
			Name:                     "map-roles",
			Type:                     discordgo.ChatApplicationCommand,
			Description:              "Map Discord to Roddbot roles",
			DefaultMemberPermissions: &admin,
			DMPermission:             &f,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "admin-role",
					Type:        discordgo.ApplicationCommandOptionMentionable,
					Description: "Which Discord role should grant the Roddbot admin role. If none, guild admins become roddbot admins.",
				},
				{
					Name:        "player-role",
					Type:        discordgo.ApplicationCommandOptionMentionable,
					Description: "Which Discord role should grant the Roddbot player role. If none, everyone is granted permission",
				},
			},
		},
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
		{
			Name:         "events",
			Type:         discordgo.ChatApplicationCommand,
			Description:  "events",
			DMPermission: &f,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "new",
					Description: "Create a new event",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "title",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Title",
							Required:    true,
						},
						{
							Name:        "start",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Start date and time (2006-01-02 15:04)",
							Required:    true,
						},
						{
							Name:        "end",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "End date and time (2006-01-02 15:04)",
							Required:    true,
						},
						{
							Name:        "location",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Location",
							Required:    true,
						},
						{
							Name:        "description",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Description",
							Required:    true,
						},
						{
							Name:        "announcement-channel",
							Type:        discordgo.ApplicationCommandOptionChannel,
							Description: "Announcement channel",
							Required:    true,
						},
						{
							Name:        "announcement-time",
							Type:        discordgo.ApplicationCommandOptionString,
							Description: "Delay the announcement (2006-01-02 15:04). Leave empty if you want it announce it now.",
						},
					},
				},
				{
					Name:        "list",
					Description: "List events",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options:     []*discordgo.ApplicationCommandOption{},
				},
			},
		},
	}
}
