package auth_usecase

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"picket/src/app"
)

var ErrVersionNotValid = app.NewForbiddenError("version not valid")
var ErrTokenNotValid = app.NewRawError("token not valid", http.StatusForbidden)

func (u *usecase) Verify(ctx context.Context, token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.secretKey), nil
	})
	if err != nil {
		log.Error().Err(err).Send()
		return -1, ErrTokenNotValid
	}
	if !t.Valid {
		return -1, ErrTokenNotValid
	}
	claims, ok := t.Claims.(*AuthClaims)
	if !ok {
		return -1, ErrTokenNotValid
	}
	user, err := u.repository.FindById(ctx, claims.UserId)
	if err != nil {
		return -1, err
	}
	if claims.Version != user.Version {
		return -1, ErrVersionNotValid
	}
	userId := claims.UserId
	return userId, nil
}
