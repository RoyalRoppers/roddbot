package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvKeyToken = "RB_DISCORD_BOT_TOKEN"
)

type Config struct {
	DiscordBotToken string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	discordToken := os.Getenv(EnvKeyToken)

	return &Config{
		DiscordBotToken: discordToken,
	}, nil
}
