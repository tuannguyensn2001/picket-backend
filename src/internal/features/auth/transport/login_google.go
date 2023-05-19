package auth_transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picket/src/app"
	"picket/src/internal/dto"
)

func (t *transport) LoginGoogle(ctx *gin.Context) {
	var input dto.LoginGoogleInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err.Error()))
	}

	result, err := t.usecase.LoginGoogle(ctx, input.Code)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
