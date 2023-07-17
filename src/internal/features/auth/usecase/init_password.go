package auth_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"picket/src/app"
	"picket/src/internal/dto"
)

func (u *usecase) InitPassword(ctx context.Context, input dto.InitPasswordInput) error {
	user, err := u.repository.FindById(ctx, input.UserId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	if len(user.Password) > 0 {
		return app.NewBadRequestError("user has password")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	user.Password = string(hashPassword)
	err = u.repository.Save(ctx, user)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}
