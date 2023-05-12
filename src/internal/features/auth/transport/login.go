package auth_transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picket/src/app"
	"picket/src/internal/dto"
)

func (t *transport) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err.Error()))
	}
	result, err := t.usecase.Login(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
