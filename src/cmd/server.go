package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"picket/src/config"
	"picket/src/internal/routes"
)

func server(config config.Config) *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()

			r.GET("/health", func(context *gin.Context) {
				context.JSON(200, gin.H{
					"message": "pong",
				})
			})

			routes.Routes(r, config)

			r.Run(fmt.Sprintf(":%s", config.Port))
		},
	}
}
