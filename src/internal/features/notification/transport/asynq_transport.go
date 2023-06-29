package notification_transport

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"picket/src/internal/dto"
)

func (t *transport) CreateNotificationWhenUserRegister(ctx context.Context, task *asynq.Task) error {
	var payload dto.NewUserRegisterPayload
	if err := json.NewDecoder(bytes.NewReader(task.Payload())).Decode(&payload); err != nil {
		return err
	}

	if err := t.usecase.CreateNotificationWhenUserRegisterSuccess(ctx, payload.UserId); err != nil {
		return err
	}

	return nil
}
