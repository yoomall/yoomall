package global

import (
	"yoomall/yoo/config"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	return config.Config
}
