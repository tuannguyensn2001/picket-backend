package config

import (
	"github.com/rs/zerolog/log"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Db                       *gorm.DB
	Port                     string
	SecretKey                string
	DbTest                   *gorm.DB
	Oauth2GoogleClientId     string
	Oauth2GoogleClientSecret string
	ClientUrl                string
	KafkaAddress             string
}

func GetConfig() (*Config, error) {
	structure, err := bootstrap()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(structure.DatabaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Error().Err(err).Send()
	}
	dbTest, _ := gorm.Open(postgres.Open(structure.DatabaseTestUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	config := &Config{
		Db:                       db,
		Port:                     structure.AppPort,
		SecretKey:                structure.AppSecretKey,
		DbTest:                   dbTest,
		Oauth2GoogleClientId:     structure.Oauth2GoogleClientId,
		Oauth2GoogleClientSecret: structure.Oauth2GoogleClientSecret,
		ClientUrl:                structure.ClientUrl,
		KafkaAddress:             structure.KafkaAddress,
	}

	return config, nil

}
