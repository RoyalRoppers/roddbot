package main

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/movitz-s/roddbot/internal/bot"
	"github.com/movitz-s/roddbot/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrateFS embed.FS

func main() {
	err := realMain()
	if err != nil {
		fmt.Println("main errored:", err.Error())
		os.Exit(1)
	}
}

func realMain() error {
	log, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer log.Sync()

	log.Info("main started")

	conf, err := config.Load()
	if err != nil {
		log.Error("could not load config", zap.Error(err))
		return err
	}
	log.Info("config loaded")

	// migrate

	db, err := sql.Open("postgres", conf.DBStr)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	d, err := iofs.New(migrateFS, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println(err.Error())
	}

	bot, err := bot.New(conf, log, db)
	if err != nil {
		log.Error("could not create bot", zap.Error(err))
		return err
	}

	err = bot.Open()
	if err != nil {
		return err
	}

	c := make(chan struct{})
	<-c

	return nil
}
