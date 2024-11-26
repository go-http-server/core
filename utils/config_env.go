package utils

import (
	"time"

	"github.com/spf13/viper"
)

type EnviromentVariables struct {
	DB_SOURCE           string        `mapstructure:"DB_SOURCE"`
	ENVIRONMENT         string        `mapstructure:"ENVIRONMENT"`
	HTTP_SERVER_ADDRESS string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TIME_EXPIRED_TOKEN  time.Duration `mapstructure:"TIME_EXPIRED_TOKEN"`
}

func LoadEnviromentVariables(path string) (config EnviromentVariables, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env") // json, xml, ...

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
