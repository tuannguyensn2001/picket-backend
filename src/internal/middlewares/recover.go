package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"picket/src/app"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			err, ok := err.(error)
			if !ok {
				log.Error().Interface("err", err).Msg("server has error")
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
				})
				return
			}

			if val, ok := err.(*app.Error); ok {
				ctx.AbortWithStatusJSON(val.Code, gin.H{
					"message": val.Message,
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
	}()
	ctx.Next()
}
