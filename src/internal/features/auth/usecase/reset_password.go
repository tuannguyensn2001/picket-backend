package auth_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"picket/src/app"
	"picket/src/internal/dto"
)

func (u *usecase) ResetPassword(ctx context.Context, input dto.ResetPasswordInput) error {
	user, err := u.repository.FindById(ctx, input.UserId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword))
	if err != nil {
		return app.NewBadRequestError("bad request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	user.Password = string(hashedPassword)
	if err := u.repository.Save(ctx, user); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil

}
