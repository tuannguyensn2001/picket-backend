package notification_repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func (r *repo) FindByToUser(ctx context.Context, userId int) ([]entities.Notification, error) {
	var notifications []entities.Notification
	db := r.GetDB(ctx).Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)})
	if err := db.WithContext(ctx).Where("\"to\" = ?", userId).Find(&notifications).Error; err != nil {
		return nil, err
	}

	m := make(map[int]bool)
	for _, item := range notifications {
		m[item.From] = true
		m[item.To] = true
	}

	listUserIds := make([]int, 0)
	for key := range m {
		listUserIds = append(listUserIds, key)
	}

	var users []entities.User
	if err := db.WithContext(ctx).Preload("Profile").Where("id in ?", listUserIds).Find(&users).Error; err != nil {
		return nil, err
	}
	mUser := make(map[int]entities.User)
	for _, user := range users {
		mUser[user.Id] = user
	}

	for i := 0; i < len(notifications); i++ {
		to := mUser[notifications[i].To]
		from := mUser[notifications[i].From]
		notifications[i].ToUser = &to
		notifications[i].FromUser = &from
	}

	return notifications, nil
}
