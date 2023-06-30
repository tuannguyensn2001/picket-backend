package routes

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"net/http"
	"picket/src/config"
	auth_repository "picket/src/internal/features/auth/repository"
	auth_transport "picket/src/internal/features/auth/transport"
	auth_usecase "picket/src/internal/features/auth/usecase"
	notification_repository "picket/src/internal/features/notification/repository"
	notification_transport "picket/src/internal/features/notification/transport"
	notification_usecase "picket/src/internal/features/notification/usecase"
	"picket/src/internal/middlewares"
	oauth2_service "picket/src/internal/services/oauth2"
)

func Routes(r *gin.Engine, config config.Config) {
	r.Use(middlewares.Recover)
	r.Use(middlewares.Cors)
	r.Use(otelgin.Middleware("picket-backend"))

	oauth2Service := oauth2_service.New(config)

	authRepository := auth_repository.New(config.Db)
	authUsecase := auth_usecase.New(authRepository, config.SecretKey, oauth2Service, config.KafkaAddress, config.Asynq)
	authTransport := auth_transport.New(authUsecase)

	notificationRepository := notification_repository.New(config.Db)
	notificationUsecase := notification_usecase.New(notificationRepository, authUsecase)
	notificationTransport := notification_transport.New(notificationUsecase)

	g := r.Group("/api")
	g.POST("/v1/auth/register", authTransport.Register)
	g.POST("/v1/auth/login", authTransport.Login)
	g.POST("/v1/auth/login/google", authTransport.LoginGoogle)
	g.GET("/v1/auth/me", middlewares.Auth(authUsecase), authTransport.GetMe)

	g.GET("/v1/notifications/own/unread/count", middlewares.Auth(authUsecase), notificationTransport.CountUnread)
	g.GET("/v1/notifications/own", middlewares.Auth(authUsecase), notificationTransport.GetOwn)

	r.POST("/google", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"access_token": "123",
		})
	})

}
