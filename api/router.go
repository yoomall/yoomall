package api

import (
	"embed"
	"html/template"
	"net/http"
	"yoomall/cmd/http/api"
	"yoomall/config"
	"yoomall/core/driver"
	httpserver "yoomall/core/http"

	"github.com/spf13/viper"

	_template "yoomall/core/template"
)

var (
	//go:embed config.yaml
	configBytes []byte

	//go:embed templates
	templateFs embed.FS
)

var conf *viper.Viper

var server httpserver.HttpServer

func init() {
	conf = config.NewConfigFromBytes(configBytes)
	server = api.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")))

	// 设置模板
	temp := template.New("main").Funcs(_template.Funcs)
	html := template.Must(_template.ParseGlobEmbedFS(temp, templateFs, "templates", "*.html"))
	server.Engine.SetHTMLTemplate(html)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server.Engine.ServeHTTP(w, r)
}
