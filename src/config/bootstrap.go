package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"picket/src/utils"
)

type Bootstrap struct {
	DatabaseUrl     string `mapstructure:"DATABASE_URL"`
	AppPort         string `mapstructure:"APP_PORT"`
	AppSecretKey    string `mapstructure:"APP_SECRET_KEY"`
	DatabaseTestUrl string `mapstructure:"DATABASE_TEST_URL"`
}

func bootstrap() (Bootstrap, error) {
	// load file .env to bootstrap struct
	// return bootstrap struct
	path, err := utils.GetGoModPath()
	if err != nil {
		return Bootstrap{}, err
	}
	path = filepath.Dir(path)
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return Bootstrap{}, err
	}
	var bootstrap Bootstrap
	if err := viper.Unmarshal(&bootstrap); err != nil {
		return Bootstrap{}, err
	}
	return bootstrap, nil
}
