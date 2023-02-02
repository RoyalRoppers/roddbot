package bot

import (
	"database/sql"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/config"
	"github.com/movitz-s/roddbot/internal/permissions"
	"go.uber.org/zap"
)

type bot struct {
	sess *discordgo.Session
	conf *config.Config
	log  *zap.Logger
	db   *sql.DB
	perm *permissions.Service

	sync.Mutex
}

func New(conf *config.Config, log *zap.Logger, db *sql.DB) (*bot, error) {
	sess, err := discordgo.New("Bot " + conf.DiscordBotToken)
	if err != nil {
		return nil, err
	}

	b := &bot{
		sess: sess,
		conf: conf,
		log:  log.Named("bot"),
		db:   db,
		perm: permissions.New(db, log),
	}

	b.sess.ShouldRetryOnRateLimit = false
	b.sess.AddHandler(b.msgCreateHandler)

	b.sess.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		b.log.Info("bot ready", zap.String("bot-name", b.sess.State.User.Username), zap.String("bot-id", b.sess.State.User.ID))

		for _, ac := range cmds() {
			cmd, err := b.sess.ApplicationCommandCreate(sess.State.User.ID, "", ac)
			if err != nil {
				log.Error("could not create cmd", zap.Error(err), zap.String("cmdname", ac.Name))
				continue
			}
			b.log.Info("cmd created", zap.String("cmd-name", ac.Name), zap.String("cmd-id", cmd.ID))
		}
	})

	return b, nil
}

func (b *bot) Open() error {
	return b.sess.Open()
}
