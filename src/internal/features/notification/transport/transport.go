package notification_transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"picket/src/utils"
)

type IUsecase interface {
	CreateNotificationWhenUserRegisterSuccess(ctx context.Context, userId int) error
	CountUnreadByUser(ctx context.Context, userId int) (int64, error)
}

type transport struct {
	usecase IUsecase
}

func New(usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) CountUnread(ctx *gin.Context) {
	userId, err := utils.GetAuthFromContext(ctx)
	if err != nil {
		panic(err)
	}
	result, err := t.usecase.CountUnreadByUser(ctx.Request.Context(), userId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "success",
	})
}
