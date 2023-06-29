package auth_usecase

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"picket/src/app"
	"picket/src/internal/dto"
	"picket/src/internal/entities"
)

func (u *usecase) Register(ctx context.Context, input dto.RegisterInput) error {

	user, err := u.repository.FindByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Send()
		return err
	}
	if user != nil {
		return app.NewBadRequestError("user already exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	err = u.repository.Transaction(ctx, func(ctx context.Context) error {
		user = &entities.User{
			Email:    input.Email,
			Password: string(password),
			Username: input.Username,
			Wallet: &entities.Wallet{
				Balance: 0,
			},
			Profile: &entities.Profile{
				AvatarUrl: "https://static.vecteezy.com/system/resources/previews/009/734/564/original/default-avatar-profile-icon-of-social-media-user-vector.jpg",
			},
		}
		count, err := u.repository.CountAllUsers(ctx)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}
		if count == 0 {
			user.IsAdmin = true
		}
		err = u.repository.Create(ctx, user)

		if err != nil {
			log.Error().Err(err).Send()
			return err
		}

		err = u.PushJobAfterRegister(ctx, user.Id)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}

		return nil
	})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil

}
