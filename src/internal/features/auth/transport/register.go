package auth_transport

import (
	"github.com/gin-gonic/gin"
	"picket/src/app"
	"picket/src/internal/dto"
)

func (t *transport) Register(ctx *gin.Context) {
	var input dto.RegisterInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err.Error()))
	}

	err := t.usecase.Register(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
