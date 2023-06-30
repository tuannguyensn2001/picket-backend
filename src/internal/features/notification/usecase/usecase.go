package notification_usecase

import (
	"context"
	"picket/src/internal/entities"
)

type IRepository interface {
	Create(ctx context.Context, notification *entities.Notification) error
	CountUnreadByUser(ctx context.Context, userId int) (int64, error)
	FindByToUser(ctx context.Context, userId int) ([]entities.Notification, error)
}

type IAuthUsecase interface {
	GetAdmin(ctx context.Context) (*entities.User, error)
	GetById(ctx context.Context, id int) (*entities.User, error)
}

type usecase struct {
	repository  IRepository
	authUsecase IAuthUsecase
}

func New(repository IRepository, authUsecase IAuthUsecase) *usecase {
	return &usecase{repository: repository, authUsecase: authUsecase}
}

func (u *usecase) GetByUser(ctx context.Context, userId int) ([]entities.Notification, error) {
	return u.repository.FindByToUser(ctx, userId)
}
