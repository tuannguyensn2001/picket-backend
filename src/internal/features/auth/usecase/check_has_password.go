package auth_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
)

func (u *usecase) CheckHasPassword(ctx context.Context, userId int) (bool, error) {
	user, err := u.repository.FindById(ctx, userId)
	if err != nil {
		log.Error().Err(err).Send()
		return false, err
	}
	return len(user.Password) > 0, nil
}
