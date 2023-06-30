package jobs

import "context"

type INotificationUsecase interface {
	CreateNotificationWhenUserRegisterSuccess(ctx context.Context, userId int) error
}
