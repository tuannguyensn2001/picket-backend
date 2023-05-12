package auth_usecase

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"picket/src/app"
	"picket/src/internal/dto"
	"time"
)

var ErrEmailOrPasswordIncorrect = app.NewRawError("username or password not valid", http.StatusBadRequest)

func (u *usecase) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error) {
	user, err := u.repository.FindByEmail(ctx, input.Email)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, ErrEmailOrPasswordIncorrect
	}

	claims := AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
		},
		Version: user.Version,
		UserId:  user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	claims.IssuedAt = jwt.NewNumericDate(time.Now().Add(96 * time.Hour))
	claims.Subject = "refresh"
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(u.secretKey))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return &dto.LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
