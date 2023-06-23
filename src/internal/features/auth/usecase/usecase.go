package auth_usecase

import (
	"context"
	"picket/src/base"
	"picket/src/internal/entities"
)

type IRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindById(ctx context.Context, id int) (*entities.User, error)
	base.IBaseRepository
}

type IOauth2Service interface {
	GetAccessTokenFromCode(ctx context.Context, code string) (string, error)
	GetUserProfileByAccessToken(ctx context.Context, accessToken string) (*entities.User, error)
}

type usecase struct {
	repository    IRepository
	secretKey     string
	oauth2Service IOauth2Service
}

func New(repository IRepository, secretKey string, oauth2Service IOauth2Service) *usecase {
	return &usecase{repository: repository, secretKey: secretKey, oauth2Service: oauth2Service}
}
