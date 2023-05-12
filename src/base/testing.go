package base

import (
	"context"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"picket/src/internal/entities"
)

func HookReset(ctx context.Context, db *gorm.DB) error {
	g, gCtx := errgroup.WithContext(ctx)
	models := []interface{}{entities.Wallet{}, entities.User{}}
	for _, model := range models {
		model := model
		g.Go(func() error {
			return db.WithContext(gCtx).Where("1 = 1").Delete(&model).Error
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
