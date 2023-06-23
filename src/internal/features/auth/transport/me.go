package auth_transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picket/src/utils"
)

func (t *transport) GetMe(ctx *gin.Context) {
	userId, err := utils.GetAuthFromContext(ctx)
	if err != nil {
		panic(err)
	}
	user, err := t.usecase.GetMe(ctx.Request.Context(), userId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}
