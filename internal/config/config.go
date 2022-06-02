package config

import (
	"github.com/rameshsunkara/go-rest-api-example/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var config *viper.Viper

// LoadConfig TODO: Add more profiles and obey all 12-Factor app rules
func LoadConfig(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Logger.Fatal("error occurred while parsing config file", zap.Error(err))
	}
}

func GetConfig() *viper.Viper {
	return config
}
