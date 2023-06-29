package auth_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"picket/src/internal/constant"
	"picket/src/internal/dto"
)

func (u *usecase) PushJobAfterRegister(ctx context.Context, userId int) error {

	payload := dto.NewUserRegisterPayload{
		UserId: userId,
	}
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(payload)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	task := asynq.NewTask(constant.JobSendNotificationWhenUserRegisterSuccess, b.Bytes())
	_, err = u.asynq.Enqueue(task)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}
