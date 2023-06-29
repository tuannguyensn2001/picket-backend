package auth_usecase

import (
	"context"
	"picket/src/internal/entities"
)

func (u *usecase) GetById(ctx context.Context, id int) (*entities.User, error) {
	return u.repository.FindById(ctx, id)
}
