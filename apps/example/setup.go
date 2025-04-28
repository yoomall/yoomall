package example

import (
	authmiddleware "yoomall/apps/auth/middleware"
	"yoomall/apps/example/handler"
	appHandlerV2 "yoomall/apps/example/handler/v2"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/lazyfury/pulse/plugins/upload"
	"github.com/spf13/viper"
)

func NewWireDefaultApp(config *viper.Viper, db *driver.DB,
	dtkHandler *handler.DtkHandler,
	menuHandler *handler.MenuHandler,
	jtkHandler *handler.JtkHandler,
) *DefaultApp {
	return &DefaultApp{
		Config: config,
		App: framework.NewApp("default", config, db, []framework.Handler{
			menuHandler,
			jtkHandler,
			dtkHandler,
		}).WithPlugins([]framework.IPlugin{
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
