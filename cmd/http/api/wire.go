//go:build wireinject
// +build wireinject

package api

import (
	"yoomall/modules/app"
	"yoomall/modules/auth"
	"yoomall/modules/common"
	"yoomall/modules/post"
	"yoomall/modules/views"
	"yoomall/yoo"
	"yoomall/yoo/driver"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewApp(conf *viper.Viper, db *driver.DB, setupEngine func(*gin.Engine) *gin.Engine) *yoo.HttpServer {
	wire.Build(NewHttpServer, NewDoc, app.WireSet, post.WireSet, auth.WireSet, common.WireSet, views.WireSet)
	return &yoo.HttpServer{}
}
