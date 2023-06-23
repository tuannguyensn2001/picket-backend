package auth_transport

import (
	"context"
	"picket/src/internal/dto"
	"picket/src/internal/entities"
)

type IUsecase interface {
	Register(ctx context.Context, input dto.RegisterInput) error
	Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error)
	LoginGoogle(ctx context.Context, code string) (*dto.LoginOutput, error)
	GetMe(ctx context.Context, userId int) (*entities.User, error)
}

type transport struct {
	usecase IUsecase
}

func New(usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}
