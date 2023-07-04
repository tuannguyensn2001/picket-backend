package profile_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"picket/src/internal/dto"
	"picket/src/internal/entities"
)

type IRepository interface {
	Save(ctx context.Context, profile *entities.Profile) error
	FindByUserId(ctx context.Context, userId int) (*entities.Profile, error)
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}

func (u *usecase) UpdateAvatar(ctx context.Context, input dto.UpdateAvatarInput) error {

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	user, err := u.repository.FindByUserId(ctx, input.UserId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	user.AvatarUrl = input.Url
	err = u.repository.Save(ctx, user)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}
