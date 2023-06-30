package base

import (
	"context"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"picket/src/config"
	"picket/src/internal/entities"
)

func HookReset(ctx context.Context, db *gorm.DB) error {

	g, gCtx := errgroup.WithContext(ctx)
	models := []interface{}{entities.Wallet{}, entities.User{}, entities.Notification{}}
	for _, model := range models {
		model := model
		g.Go(func() error {
			return db.WithContext(gCtx).Session(&gorm.Session{
				Logger: logger.Default.LogMode(logger.Error),
			}).Where("1 = 1").Delete(&model).Error
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

type callback = func(db *gorm.DB)

func TestWithDatabase(f callback) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	db := cfg.DbTest
	defer func() {
		err := HookReset(context.TODO(), db)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}()
	f(db)
}
