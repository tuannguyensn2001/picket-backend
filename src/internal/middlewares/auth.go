package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"picket/src/app"
	"strings"
)

type IAUthUsecase interface {
	Verify(ctx context.Context, token string) (int64, error)
}

func Auth(usecase IAUthUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if len(token) == 0 {
			panic(app.NewForbiddenError("token not valid"))
		}
		split := strings.Split(token, " ")
		if len(split) != 2 {
			panic(app.NewForbiddenError("token not valid"))
		}
		if split[0] != "Bearer" {
			panic(app.NewForbiddenError("token not valid"))
		}
		userId, err := usecase.Verify(ctx.Request.Context(), split[1])
		if err != nil {
			panic(app.NewForbiddenError("token not valid"))
		}
		ctx.Set("user_id", userId)
		ctx.Next()
	}
}
