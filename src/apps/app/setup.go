package app

import (
	"yoomall/src/apps/app/handler"
	authmiddleware "yoomall/src/apps/auth/middleware"
	"yoomall/src/core"
	"yoomall/src/core/driver"
	"yoomall/src/core/plugins/upload"

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
		App: core.NewApp("default", config, db, []core.Handler{
			menuHandler,
			jtkHandler,
			dtkHandler,
		}).WithPlugins([]core.IPlugin{
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