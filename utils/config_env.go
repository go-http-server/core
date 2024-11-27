package utils

import (
	"time"

	"github.com/spf13/viper"
)

type EnviromentVariables struct {
	DB_SOURCE             string        `mapstructure:"DB_SOURCE"`
	ENVIRONMENT           string        `mapstructure:"ENVIRONMENT"`
	HTTP_SERVER_ADDRESS   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TIME_EXPIRED_TOKEN    time.Duration `mapstructure:"TIME_EXPIRED_TOKEN"`
	EMAIL_ADDRESS_SENDER  string        `mapstructure:"EMAIL_ADDRESS_SENDER"`
	EMAIL_PASSWORD_SENDER string        `mapstructure:"EMAIL_PASSWORD_SENDER"`
	EMAIL_USERNAME_SENDER string        `mapstructure:"EMAIL_USERNAME_SENDER"`
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
