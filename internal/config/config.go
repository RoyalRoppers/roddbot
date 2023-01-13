package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordBotToken string
	DBStr           string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DiscordBotToken: os.Getenv("RB_DISCORD_BOT_TOKEN"),
		DBStr:           os.Getenv("RB_DB_DSN"),
	}, nil
}
