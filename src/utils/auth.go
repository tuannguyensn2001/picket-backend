package utils

import (
	"context"
	"picket/src/app"
)

func GetAuthFromContext(ctx context.Context) (int, error) {
	value, ok := ctx.Value("user_id").(int)
	if !ok {
		return -1, app.NewForbiddenError("user_id not found in context")
	}
	return value, nil
}
