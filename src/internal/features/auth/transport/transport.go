package auth_transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"picket/src/internal/dto"
	"picket/src/internal/entities"
	"picket/src/utils"
)

type IUsecase interface {
	Register(ctx context.Context, input dto.RegisterInput) error
	Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error)
	LoginGoogle(ctx context.Context, code string) (*dto.LoginOutput, error)
	GetMe(ctx context.Context, userId int) (*entities.User, error)
	CheckHasPassword(ctx context.Context, userId int) (bool, error)
	InitPassword(ctx context.Context, input dto.InitPasswordInput) error
	ResetPassword(ctx context.Context, input dto.ResetPasswordInput) error
}

type transport struct {
	usecase IUsecase
}

func New(usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) CheckHasPassword(ctx *gin.Context) {
	userId, err := utils.GetAuthFromContext(ctx)
	if err != nil {
		panic(err)
	}
	result, err := t.usecase.CheckHasPassword(ctx.Request.Context(), userId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "success",
	})
}

func (t *transport) InitPassword(ctx *gin.Context) {
	userId, err := utils.GetAuthFromContext(ctx)
	if err != nil {
		panic(err)
	}

	var input dto.InitPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		panic(err)
	}
	input.UserId = userId

	if err := t.usecase.InitPassword(ctx.Request.Context(), input); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (t *transport) ResetPassword(ctx *gin.Context) {
	userId, err := utils.GetAuthFromContext(ctx)
	if err != nil {
		panic(err)

	}

	var input dto.ResetPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		panic(err)
	}
	input.UserId = userId

	if err := t.usecase.ResetPassword(ctx.Request.Context(), input); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
