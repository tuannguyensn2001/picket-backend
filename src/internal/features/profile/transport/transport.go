package profile_transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"picket/src/app"
	"picket/src/internal/dto"
	"picket/src/utils"
)

type IUsecase interface {
	UpdateAvatar(ctx context.Context, input dto.UpdateAvatarInput) error
}

type transport struct {
	usecase IUsecase
}

func New(usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) UpdateAvatar(ctx *gin.Context) {
	userId, err := utils.GetAuthFromContext(ctx)
	if err != nil {
		panic(err)
	}

	var input dto.UpdateAvatarInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		panic(app.NewBadRequestError(err.Error()))
	}
	input.UserId = userId

	err = t.usecase.UpdateAvatar(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
