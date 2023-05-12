package wallet_repository

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

func (r *repo) FindByUserId(ctx context.Context, userId int) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := r.GetDB(ctx).WithContext(ctx).Where("user_id = ?", userId).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *repo) Save(ctx context.Context, wallet *entities.Wallet) error {
	return r.GetDB(ctx).WithContext(ctx).Save(wallet).Error
}
