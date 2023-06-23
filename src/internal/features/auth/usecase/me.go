package auth_usecase

import (
	"context"
	"picket/src/internal/entities"
)

func (u *usecase) GetMe(ctx context.Context, userId int) (*entities.User, error) {
	return u.repository.FindById(ctx, userId)
}
