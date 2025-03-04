package utils

import (
	"time"

	"github.com/spf13/viper"
)

type EnvironmentVariables struct {
	DB_SOURCE             string        `mapstructure:"DB_SOURCE"`
	ENVIRONMENT           string        `mapstructure:"ENVIRONMENT"`
	HTTP_SERVER_ADDRESS   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TIME_EXPIRED_TOKEN    time.Duration `mapstructure:"TIME_EXPIRED_TOKEN"`
	EMAIL_ADDRESS_SENDER  string        `mapstructure:"EMAIL_ADDRESS_SENDER"`
	EMAIL_PASSWORD_SENDER string        `mapstructure:"EMAIL_PASSWORD_SENDER"`
	EMAIL_USERNAME_SENDER string        `mapstructure:"EMAIL_USERNAME_SENDER"`
	REDIS_ADDRESS_SERVER  string        `mapstructure:"REDIS_ADDRESS_SERVER"`
	REDIS_PASSWORD_SERVER string        `mapstructure:"REDIS_PASSWORD_SERVER"`
	TELEGRAM_BOT_TOKEN    string        `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TELEGRAM_CHAT_ID      string        `mapstructure:"TELEGRAM_CHAT_ID"`
}

func LoadEnviromentVariables(path string) (config EnvironmentVariables, err error) {
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
