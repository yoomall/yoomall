//go:build wireinject
// +build wireinject

package server

import (
	"yoomall/modules/app"
	"yoomall/modules/auth"
	"yoomall/modules/common"
	"yoomall/modules/post"
	"yoomall/modules/views"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

func NewApp(conf *viper.Viper, db *driver.DB, setupEngine func(*gin.Engine) *gin.Engine) *framework.HttpServer {
	wire.Build(NewHttpServer,
		NewDoc,
		app.WireSet,
		app.V2WireSet,
		post.WireSet, auth.WireSet, common.WireSet, views.WireSet,
	)
	return &framework.HttpServer{}
}
