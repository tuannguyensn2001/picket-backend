package notification_repository

import (
	"context"
	"gorm.io/gorm"
	"picket/src/base"
	"picket/src/internal/entities"
)

type repo struct {
	base.Repository
}

func New(db *gorm.DB) *repo {
	return &repo{
		base.Repository{
			Db: db,
		},
	}
}

func (r *repo) Create(ctx context.Context, notification *entities.Notification) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(notification).Error
}

func (r *repo) CountUnreadByUser(ctx context.Context, userId int) (int64, error) {
	var count int64
	ctx = context.WithValue(ctx, "log-db", true)
	db := r.GetDB(ctx)
	if err := db.WithContext(ctx).Model(&entities.Notification{}).Where("\"to\" = ? and read_at = ?", userId, 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
