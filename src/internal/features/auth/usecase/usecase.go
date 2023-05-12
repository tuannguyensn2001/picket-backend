package auth_usecase

import (
	"context"
	"picket/src/internal/entities"
)

type IRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindById(ctx context.Context, id int) (*entities.User, error)
}

type usecase struct {
	repository IRepository
	secretKey  string
}

func New(repository IRepository, secretKey string) *usecase {
	return &usecase{repository: repository, secretKey: secretKey}
}
