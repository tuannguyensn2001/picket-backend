package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"net/http"
	"picket/src/config"
	auth_repository "picket/src/internal/features/auth/repository"
	auth_transport "picket/src/internal/features/auth/transport"
	auth_usecase "picket/src/internal/features/auth/usecase"
	notification_repository "picket/src/internal/features/notification/repository"
	notification_transport "picket/src/internal/features/notification/transport"
	notification_usecase "picket/src/internal/features/notification/usecase"
	profile_repository "picket/src/internal/features/profile/repository"
	profile_transport "picket/src/internal/features/profile/transport"
	profile_usecase "picket/src/internal/features/profile/usecase"
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

	profileRepositroy := profile_repository.New(config.Db)
	profileUsecase := profile_usecase.New(profileRepositroy)
	profileTransport := profile_transport.New(profileUsecase)

	checkAuth := middlewares.Auth(authUsecase)

	g := r.Group("/api")
	g.POST("/v1/auth/register", authTransport.Register)
	g.POST("/v1/auth/login", authTransport.Login)
	g.POST("/v1/auth/login/google", authTransport.LoginGoogle)
	g.GET("/v1/auth/me", checkAuth, authTransport.GetMe)
	g.GET("/v1/auth/has-password", checkAuth, authTransport.CheckHasPassword)
	g.PUT("/v1/auth/init-password", checkAuth, authTransport.InitPassword)
	g.PUT("/v1/auth/reset-password", checkAuth, authTransport.ResetPassword)

	g.PUT("/v1/profiles/avatar", checkAuth, profileTransport.UpdateAvatar)

	g.GET("/v1/notifications/own/unread/count", checkAuth, notificationTransport.CountUnread)
	g.GET("/v1/notifications/own", checkAuth, notificationTransport.GetOwn)

	r.POST("/api/v1/uploads", checkAuth, func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			panic(err)
		}

		f, err := file.Open()
		if err != nil {
			panic(err)
		}

		fileName := uuid.New()
		_, err = config.Minio.PutObject(context.Request.Context(), "picket", fileName.String(), f, file.Size, minio.PutObjectOptions{})
		if err != nil {
			panic(err)
		}

		url := fmt.Sprintf("%s/%s/%s", config.Minio.EndpointURL().String(), "picket", fileName.String())

		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    url,
		})

	})

	r.POST("/google", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"access_token": "123",
		})
	})

}
