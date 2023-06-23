package auth_usecase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"picket/src/internal/dto"
	"picket/src/internal/entities"
	"time"
)

func (u *usecase) LoginGoogle(ctx context.Context, code string) (*dto.LoginOutput, error) {
	result, err := u.oauth2Service.GetAccessTokenFromCode(ctx, code)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	googleAccount, err := u.oauth2Service.GetUserProfileByAccessToken(ctx, result)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	user, err := u.repository.FindByEmail(ctx, googleAccount.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Send()
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &entities.User{
			Email:    googleAccount.Email,
			Username: googleAccount.Username,
			Profile: &entities.Profile{
				AvatarUrl: googleAccount.Profile.AvatarUrl,
			},
		}
		err = u.repository.Create(ctx, user)
		if err != nil {
			log.Error().Err(err).Send()
			return nil, err
		}
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
