package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
	"lazyfury.github.com/yoomall-server/libs/dtk"
)

type Config struct {
	MySQL struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	} `default:"{}" yaml:"mysql"`

	PG struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	} `default:"{}" yaml:"postgres"`

	DEBUG bool `default:"false"`

	HTTP struct {
		Port int
	}

	DTK *dtk.DtkConfig `default:"{}" yaml:"dtk"`

	STORAGT_PATH string `yaml:"storage_path"`
}

func GetConfig(file string) *Config {
	// 读取配置
	b, err := os.ReadFile(file)
	if err != nil {
		genConfigFile(file)
		panic("读取配置错误: " + err.Error())
	}
	// yaml
	var config Config
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		panic("解析配置错误: " + err.Error())
	}
	return &config
}

func genConfigFile(path string) {
	var defConfig *Config = &Config{
		MySQL: struct {
			Host     string
			Port     int
			Username string
			Password string
			Database string
		}{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "123456",
			Database: "yoomall",
		},

		PG: struct {
			Host     string
			Port     int
			Username string
			Password string
			Database string
		}{
			Host:     "127.0.0.1",
			Port:     5432,
			Username: "postgres",
			Password: "123456",
			Database: "yoomall",
		},

		DEBUG: false,
		HTTP: struct {
			Port int
		}{
			Port: 8900,
		},

		DTK: &dtk.DtkConfig{
			AppKey:    "app_key",
			AppSecret: "app_secret",
			AppUrl:    "https://api.dtkmall.com",
		},
		STORAGT_PATH: "./storage",
	}
	// 生成配置
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	str, err := yaml.Marshal(defConfig)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(str)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func (c *Config) MysqlDsn() string {
	portStr := strconv.Itoa(c.MySQL.Port)
	return c.MySQL.Username + ":" + c.MySQL.Password + "@tcp(" + c.MySQL.Host + ":" + portStr + ")/" + c.MySQL.Database
}

func (c *Config) PgDsn() string {
	portStr := strconv.Itoa(c.PG.Port)
	template := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	return fmt.Sprintf(template, c.PG.Host, portStr, c.PG.Username, c.PG.Password, c.PG.Database)
}
