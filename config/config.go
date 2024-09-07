package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core/constants"
)

func NewConfig() *viper.Viper {
	config_file_path := os.Getenv(constants.YOO_CONFIG)
	if config_file_path != "" {
		return GetConfig(config_file_path)
	}
	return GetConfig("./config.yaml")
}

func GetConfig(file string) *viper.Viper {
	config := viper.New()

	config.SetConfigFile(file)
	if err := config.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file, err: ", err)
		os.Exit(1)
	}
	return config
}

type MysqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func GetMysqlConfig(config *viper.Viper) *MysqlConfig {
	return &MysqlConfig{
		Host:     config.GetString("mysql.host"),
		Port:     config.GetInt("mysql.port"),
		Username: config.GetString("mysql.username"),
		Password: config.GetString("mysql.password"),
		DbName:   config.GetString("mysql.db_name"),
	}
}

type HttpConfig struct {
	Port int
}

func GetHttpConfig(config *viper.Viper) *HttpConfig {
	return &HttpConfig{
		Port: config.GetInt("http.port"),
	}
}
