package initalizers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBLocal         string `mapstructure:"DB_LOCAL"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	DBName          string `mapstructure:"DB_NAME"`
	DBHostPort      string `mapstructure:"DB_HOST_PORT"`
	DBContainerPort string `mapstructure:"DB_CONTAINER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	return
}
