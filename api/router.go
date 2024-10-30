package api

import (
	"net/http"
	"yoomall/cmd/http/api"
	"yoomall/config"
	"yoomall/core/driver"
	httpserver "yoomall/core/http"

	"github.com/spf13/viper"

	_ "embed"
)

var (
	//go:embed config.yaml
	configBytes []byte
)

var conf *viper.Viper

var server httpserver.HttpServer

func init() {
	conf = config.NewConfigFromBytes(configBytes)
	config.Init(true, conf)
	server = api.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server.Engine.ServeHTTP(w, r)
}
