package server

import (
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

func writeDefaultConfig(configs ...map[string]interface{}) error {
	config := viper.New()
	config.SetConfigFile("preview-config.yaml")
	for _, m := range configs {
		for k, v := range m {
			config.Set(k, v)
		}
	}
	return config.WriteConfig()
}

func previewDefaultConfig() error {
	return writeDefaultConfig(map[string]interface{}{
		"mysql": driver.MysqlConfig{},
		"DEBUG": true,
	})
}
