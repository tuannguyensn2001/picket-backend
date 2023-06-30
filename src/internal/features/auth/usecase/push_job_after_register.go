package auth_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"picket/src/internal/jobs"
)

func (u *usecase) PushJobAfterRegister(ctx context.Context, userId int) error {

	task, err := jobs.NewUserRegisterSuccessJob(userId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	_, err = u.asynq.EnqueueContext(ctx, task)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}
