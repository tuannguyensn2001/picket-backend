package notification_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"picket/src/internal/constant"
	"picket/src/internal/entities"
)

func (u *usecase) CreateNotificationWhenUserRegisterSuccess(ctx context.Context, userId int) error {
	user, err := u.authUsecase.GetById(ctx, userId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	admin, err := u.authUsecase.GetAdmin(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	notification := entities.Notification{
		From:     admin.Id,
		To:       user.Id,
		Type:     constant.WELCOME_TYPE,
		Template: constant.WELCOME_TEMPLATE,
		Payload: entities.Payload{
			"name": user.Username,
		},
	}
	err = u.repository.Create(ctx, &notification)
	if err != nil {
		log.Error().Err(err).Send()
	}

	return nil

}
