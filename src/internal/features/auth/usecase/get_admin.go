package auth_usecase

import (
	"context"
	"picket/src/internal/entities"
)

func (u *usecase) GetAdmin(ctx context.Context) (*entities.User, error) {
	return u.repository.FindAdmin(ctx)
}
