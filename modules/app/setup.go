package app

import (
	"yoomall/modules/app/handler"
	appHandlerV2 "yoomall/modules/app/handler/v2"
	authmiddleware "yoomall/modules/auth/middleware"
	yoo "yoomall/yoo"
	"yoomall/yoo/driver"
	"yoomall/yoo/plugins/upload"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewWireDefaultApp(config *viper.Viper, db *driver.DB,
	dtkHandler *handler.DtkHandler,
	menuHandler *handler.MenuHandler,
	jtkHandler *handler.JtkHandler,
) *DefaultApp {
	return &DefaultApp{
		Config: config,
		App: yoo.NewApp("default", config, db, []yoo.Handler{
			menuHandler,
			jtkHandler,
			dtkHandler,
		}).WithPlugins([]yoo.IPlugin{
			upload.NewUploadPlugin().WithMiddlewares([]gin.HandlerFunc{
				authmiddleware.NewAuthMiddleware(db, true, false),
			}),
		}),
	}
}

var handlerSet = wire.NewSet(
	handler.NewDtkHandler,
	handler.NewMenuHandler,
	handler.NewJtkHandler,
)

var WireSet = wire.NewSet(NewWireDefaultApp, handlerSet)

var V2WireSet = wire.NewSet(NewDefaultV2App, appHandlerV2.NewBaseHandler)
