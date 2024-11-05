package config

import (
	"api/yoo/global"
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	return NewConfigFromFile()
}

func NewConfigFromBytes(_bytes []byte) *viper.Viper {
	if len(_bytes) == 0 {
		panic("config file bytes is empty")
	}

	config := viper.New()
	config.SetConfigType("yaml")

	if err := config.ReadConfig(bytes.NewBuffer(_bytes)); err != nil {
		panic("failed to read config file, err: " + err.Error())
	}
	setup(config)
	return config
}

func NewConfigFromFile() *viper.Viper {

	config := viper.New()
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	config.AddConfigPath("./config/")
	config.AddConfigPath("$HOME/.yoomall/")
	config.SetConfigName("config")

	config.SetEnvPrefix("yoomall")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file, err: ", err)
		// gen default config.yaml
		config.WriteConfigAs("./config.yaml")
		os.Exit(1)
	}
	setup(config)
	return config
}

func setup(config *viper.Viper) {
	global.Config = config
}
