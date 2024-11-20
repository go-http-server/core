package utils

import (
	"github.com/spf13/viper"
)

type EnviromentVariables struct {
	DB_SOURCE string `mapstructure:"DB_SOURCE"`
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
