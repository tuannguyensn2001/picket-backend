package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Db        *gorm.DB
	Port      string
	SecretKey string
	DbTest    *gorm.DB
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
	dbTest, _ := gorm.Open(postgres.Open(structure.DatabaseTestUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	config := &Config{
		Db:        db,
		Port:      structure.AppPort,
		SecretKey: structure.AppSecretKey,
		DbTest:    dbTest,
	}

	return config, nil

}
