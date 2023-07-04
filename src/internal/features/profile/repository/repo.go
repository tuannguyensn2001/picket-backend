package profile_repository

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

func (r *repo) Save(ctx context.Context, profile *entities.Profile) error {
	return r.GetDB(ctx).WithContext(ctx).Save(profile).Error
}

func (r *repo) FindByUserId(ctx context.Context, userId int) (*entities.Profile, error) {
	var profile entities.Profile
	err := r.GetDB(ctx).WithContext(ctx).Where("user_id = ?", userId).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
