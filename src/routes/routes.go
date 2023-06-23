package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picket/src/config"
	auth_repository "picket/src/internal/features/auth/repository"
	auth_transport "picket/src/internal/features/auth/transport"
	auth_usecase "picket/src/internal/features/auth/usecase"
	"picket/src/internal/middlewares"
	oauth2_service "picket/src/internal/services/oauth2"
)

func Routes(r *gin.Engine, config config.Config) {
	r.Use(middlewares.Recover)
	r.Use(middlewares.Cors)

	oauth2Service := oauth2_service.New(config)

	authRepository := auth_repository.New(config.Db)
	authUsecase := auth_usecase.New(authRepository, config.SecretKey, oauth2Service)
	authTransport := auth_transport.New(authUsecase)

	g := r.Group("/api")
	g.POST("/v1/auth/register", authTransport.Register)
	g.POST("/v1/auth/login", authTransport.Login)
	g.POST("/v1/auth/login/google", authTransport.LoginGoogle)

	r.POST("/google", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"access_token": "123",
		})
	})
}
