package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"picket/src/config"
	"picket/src/routes"
)

func server(config config.Config) *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()
			mux := asynq.NewServeMux()

			r.GET("/health", func(context *gin.Context) {
				context.JSON(200, gin.H{
					"message": "pong",
				})
			})

			routes.Routes(r, config, mux)

			go func() {
				if err := config.AsynqServer.Run(mux); err != nil {
					log.Error().Err(err).Send()
				}
			}()

			err := r.Run(fmt.Sprintf(":%s", config.Port))
			if err != nil {
				return
			}
		},
	}
}
