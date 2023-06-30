package jobs

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"picket/src/internal/constant"
	"picket/src/internal/dto"
)

func NewUserRegisterSuccessJob(userId int) (*asynq.Task, error) {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(dto.NewUserRegisterPayload{
		UserId: userId,
	}); err != nil {
		return nil, err
	}
	return asynq.NewTask(constant.NewUserRegisterSuccessJob, b.Bytes()), nil
}

type newUserRegisterSuccessProcessor struct {
	notificationUsecase INotificationUsecase
}

func NewUserRegisterSuccessProcessor(notificationUsecase INotificationUsecase) *newUserRegisterSuccessProcessor {
	return &newUserRegisterSuccessProcessor{
		notificationUsecase: notificationUsecase,
	}
}

func (p *newUserRegisterSuccessProcessor) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload dto.NewUserRegisterSuccessJobPayload
	if err := json.NewDecoder(bytes.NewReader(task.Payload())).Decode(&payload); err != nil {
		return err
	}
	log.Info().Interface("msg", payload).Send()

	if err := p.notificationUsecase.CreateNotificationWhenUserRegisterSuccess(ctx, payload.UserId); err != nil {
		return err
	}

	return nil
}
