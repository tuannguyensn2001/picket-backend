package auth_repository

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
	return &repo{base.Repository{Db: db}}
}

func (r *repo) Create(ctx context.Context, user *entities.User) error {
	return r.GetDB(ctx).WithContext(ctx).Create(user).Error
}

func (r *repo) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	return r.GetDB(ctx).WithContext(ctx).Create(profile).Error
}

func (r *repo) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var result entities.User
	if err := r.GetDB(ctx).WithContext(ctx).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repo) FindById(ctx context.Context, id int) (*entities.User, error) {
	var result entities.User
	if err := r.GetDB(ctx).WithContext(ctx).Where("id = ?", id).Preload("Profile").First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
