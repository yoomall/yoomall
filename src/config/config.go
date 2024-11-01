package config

import (
	"bytes"
	"fmt"
	"os"

	"yoomall/src/core/constants"

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
	err := config.ReadConfig(bytes.NewBuffer(_bytes))
	if err != nil {
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

var _viper *viper.Viper

var Config *struct {
	*viper.Viper
	Port     int
	DEBUG    bool
	MysqlDsn string
}

var VITE_URL string
var VITE_DEBUG bool
var VITE_BUILD_DIR string

func setup(config *viper.Viper) {

	_viper = config

	// 常用配置： viper 的用法很难收集配置，记录一些常用的配置，方便以后使用
	Config = &struct {
		*viper.Viper
		Port     int
		DEBUG    bool
		MysqlDsn string
	}{
		Viper:    _viper,
		Port:     _viper.GetInt(constants.PORT),
		DEBUG:    _viper.GetBool(constants.DEBUG),
		MysqlDsn: _viper.GetString(constants.MYSQL_DSN),
	}

	VITE_URL = _viper.GetString(constants.VITE_URL)
	VITE_DEBUG = _viper.GetBool(constants.VITE_DEBUG)
	VITE_BUILD_DIR = _viper.GetString(constants.VITE_BUILD_DIR)
}
