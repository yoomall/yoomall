package api

import (
	"embed"
	"html/template"
	"net/http"
	"yoomall/cmd/http/api"
	"yoomall/config"
	"yoomall/core/driver"
	httpserver "yoomall/core/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_template "yoomall/core/template"
)

var (
	//go:embed config.yaml
	configBytes []byte

	//go:embed templates/**
	templateFs embed.FS

	//go:embed manifest.json
	viteManifestJSON []byte
)

var conf *viper.Viper

var server httpserver.HttpServer

func init() {
	conf = config.NewConfigFromBytes(configBytes)

	server = api.NewApp(conf, driver.NewPostgresDB(conf.GetString("postgres.dsn")), func(e *gin.Engine) *gin.Engine {
		// 设置模板
		temp := template.New("main").Funcs(_template.Funcs(viteManifestJSON))
		html := template.Must(_template.ParseGlobEmbedFS(temp, templateFs, "templates", "*.html"))
		e.SetHTMLTemplate(html)
		return e
	})

}

func Handler_api_router_go(w http.ResponseWriter, r *http.Request) {
	server.Engine.ServeHTTP(w, r)
}
