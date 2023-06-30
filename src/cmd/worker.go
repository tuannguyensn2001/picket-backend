package cmd

import (
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"picket/src/config"
	"picket/src/internal/constant"
	auth_repository "picket/src/internal/features/auth/repository"
	auth_usecase "picket/src/internal/features/auth/usecase"
	notification_repository "picket/src/internal/features/notification/repository"
	notification_usecase "picket/src/internal/features/notification/usecase"
	"picket/src/internal/jobs"
	oauth2_service "picket/src/internal/services/oauth2"
)

func worker(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use: "worker",
		Run: func(cmd *cobra.Command, args []string) {
			srv := cfg.AsynqServer
			mux := asynq.NewServeMux()

			oauth2Service := oauth2_service.New(cfg)
			authRepository := auth_repository.New(cfg.Db)
			authUsecase := auth_usecase.New(authRepository, cfg.SecretKey, oauth2Service, cfg.KafkaAddress, cfg.Asynq)
			notificationRepository := notification_repository.New(cfg.Db)
			notificationUsecase := notification_usecase.New(notificationRepository, authUsecase)

			mux.Handle(constant.NewUserRegisterSuccessJob, jobs.NewUserRegisterSuccessProcessor(notificationUsecase))

			if err := srv.Run(mux); err != nil {
				log.Fatal().Err(err).Send()
			}
		},
	}
}
