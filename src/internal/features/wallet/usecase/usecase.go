package wallet_usecase

import (
	"context"
	"picket/src/base"
	"picket/src/internal/entities"
)

type IRepository interface {
	base.IBaseRepository
	FindByUserId(ctx context.Context, userId int) (*entities.Wallet, error)
	Save(ctx context.Context, wallet *entities.Wallet) error
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}
