//go:build wireinject
// +build wireinject

package api

import (
	"yoomall/apps/app"
	"yoomall/apps/auth"
	"yoomall/apps/common"
	"yoomall/apps/post"
	"yoomall/apps/views"
	"yoomall/core/driver"
	httpserver "yoomall/core/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewApp(conf *viper.Viper, db *driver.DB, setHTMLTemplate func(*gin.Engine) *gin.Engine) httpserver.HttpServer {
	wire.Build(NewHttpServer, NewDoc, app.WireSet, post.WireSet, auth.WireSet, common.WireSet, views.WireSet)
	return httpserver.HttpServer{}
}
