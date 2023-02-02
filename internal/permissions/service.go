package permissions

import (
	"context"
	"database/sql"
	"sync"

	"github.com/movitz-s/roddbot/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type Service struct {
	db  *sql.DB
	log *zap.Logger

	mappings models.GuildRoleMappingSlice

	x sync.RWMutex
}

func New(db *sql.DB, log *zap.Logger) *Service {
	s := &Service{
		db:  db,
		log: log.Named("permissions"),
	}

	go s.init()

	return s
}

func (s *Service) init() {
	s.x.Lock()
	defer s.x.Unlock()

	s.fillCache()
}

func (s *Service) fillCache() error {
	s.log.Info("filling the cache")

	mappings, err := models.GuildRoleMappings().All(context.Background(), s.db)
	if err != nil {
		s.log.Error("could not get guild role mappings", zap.Error(err))
		return err
	}

	s.mappings = mappings
	return nil
}

func (s *Service) MapRoles(guildID string, adminRoleID, playerRoleID *string) error {
	s.x.Lock()
	defer s.x.Unlock()

	x := &models.GuildRoleMapping{
		GuildID:      guildID,
		AdminRoleID:  null.StringFromPtr(adminRoleID),
		PlayerRoleID: null.StringFromPtr(playerRoleID),
	}

	err := x.Upsert(context.TODO(), s.db, true, nil, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return s.fillCache()
}

func (s *Service) GetMappings(guildID string) *models.GuildRoleMapping {
	s.x.RLock()
	defer s.x.RUnlock()

	for _, v := range s.mappings {
		if v.GuildID == guildID {
			return v
		}
	}

	return &models.GuildRoleMapping{GuildID: guildID} // hacky, but give me a break
}
