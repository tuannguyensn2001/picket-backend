package routes

import (
	"github.com/gin-gonic/gin"
	"picket/src/config"
	auth_repository "picket/src/internal/features/auth/repository"
	auth_transport "picket/src/internal/features/auth/transport"
	auth_usecase "picket/src/internal/features/auth/usecase"
	"picket/src/internal/middlewares"
)

func Routes(r *gin.Engine, config config.Config) {
	r.Use(middlewares.Recover)

	authRepository := auth_repository.New(config.Db)
	authUsecase := auth_usecase.New(authRepository, config.SecretKey)
	authTransport := auth_transport.New(authUsecase)

	g := r.Group("/api")
	g.POST("/v1/auth/register", authTransport.Register)
	g.POST("/v1/auth/login", authTransport.Login)
}
