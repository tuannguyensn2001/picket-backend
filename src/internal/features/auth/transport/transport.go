package auth_transport

import (
	"context"
	"picket/src/internal/dto"
)

type IUsecase interface {
	Register(ctx context.Context, input dto.RegisterInput) error
	Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error)
	LoginGoogle(ctx context.Context, code string) (*dto.LoginOutput, error)
}

type transport struct {
	usecase IUsecase
}

func New(usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}
