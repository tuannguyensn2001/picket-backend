package cmd

import (
	"github.com/spf13/cobra"
	"picket/src/config"
)

func Root(config config.Config) *cobra.Command {
	server := server(config)

	root := cobra.Command{}

	root.AddCommand(server)

	return &root
}
