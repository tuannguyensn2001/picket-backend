package auth_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"picket/src/internal/entities"
)

func (u *usecase) PushToKafkaAfterCreate(ctx context.Context, user entities.User) error {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(u.kafkaAddress),
		Topic:                  "user_created",
		RequiredAcks:           kafka.RequireOne,
		AllowAutoTopicCreation: true,
		BatchSize:              1,
	}

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(user); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	ctx, span := tracer.Start(ctx, "push to kafka")
	defer span.End()
	err := w.WriteMessages(ctx, kafka.Message{
		Value: b.Bytes(),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}
