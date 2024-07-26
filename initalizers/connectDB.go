package initalizers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBContainerPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}
