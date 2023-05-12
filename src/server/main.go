package main

import (
	"github.com/rs/zerolog/log"
	"picket/src/cmd"
	config2 "picket/src/config"
)

func main() {
	config, err := config2.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("fail to load config")
	}
	root := cmd.Root(*config)

	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Msg("fail to execute root command")
	}

}
