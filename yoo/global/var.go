package global

import "github.com/spf13/viper"

var (
	// init by yoo.config.Newxxx
	Config *viper.Viper = viper.New()
)

func Init(config *viper.Viper) {
	Config = config
}
