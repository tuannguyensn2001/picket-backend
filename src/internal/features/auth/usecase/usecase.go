package auth_usecase

import (
	"context"
	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel"
	"picket/src/base"
	"picket/src/internal/entities"
)

var tracer = otel.Tracer("auth_usecase")

type IRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindById(ctx context.Context, id int) (*entities.User, error)
	base.IBaseRepository
	CountAllUsers(ctx context.Context) (int64, error)
	FindAdmin(ctx context.Context) (*entities.User, error)
	Save(ctx context.Context, user *entities.User) error
}

type IOauth2Service interface {
	GetAccessTokenFromCode(ctx context.Context, code string) (string, error)
	GetUserProfileByAccessToken(ctx context.Context, accessToken string) (*entities.User, error)
}

type usecase struct {
	repository    IRepository
	secretKey     string
	oauth2Service IOauth2Service
	kafkaAddress  string
	asynq         *asynq.Client
}

func New(repository IRepository, secretKey string, oauth2Service IOauth2Service, kafkaAddress string, asynq *asynq.Client) *usecase {
	return &usecase{repository: repository, secretKey: secretKey, oauth2Service: oauth2Service, kafkaAddress: kafkaAddress, asynq: asynq}
}
