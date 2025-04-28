package projConfig

import (
	"yoomall/libs/dtk"
	"yoomall/libs/jutuike"

	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

func writeDefaultConfig(configs ...map[string]interface{}) error {
	config := viper.New()
	config.SetConfigFile("./config/preview.config.yaml")
	for _, m := range configs {
		for k, v := range m {
			config.Set(k, v)
		}
	}
	return config.WriteConfig()
}

func PreviewDefaultConfig() error {
	return writeDefaultConfig(map[string]interface{}{
		"mysql":        driver.MysqlConfig{},
		"DEBUG":        true,
		"dtk":          dtk.DtkConfig{},
		"jtl":          jutuike.JtkConfig{},
		"storage_path": "./storage",
		"vite": map[string]interface{}{
			"url":         "",
			"public_dir":  "",
			"debug":       true,
			"assets_root": "/",
		},
		"site": map[string]interface{}{
			"title":       "",
			"description": "",
			"keywords":    "",
			"logo":        "",
			"author":      "",
			"favicon":     "",
		},
	})
}
