package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/app/handler"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

func NewWireDefaultApp(config *viper.Viper, db *driver.DB,
	dtkHandler *handler.DtkHandler,
	menuHandler *handler.MenuHandler,
	jtkHandler *handler.JtkHandler,
) *DefaultApp {
	return &DefaultApp{
		Config: config,
		AppImpl: core.NewAppImpl("default", config, db, []core.Handler{
			menuHandler,
			jtkHandler,
		}),
		AuthMiddlewares: []gin.HandlerFunc{},
	}
}

var handlerSet = wire.NewSet(
	handler.NewDtkHandler,
	handler.NewMenuHandler,
	handler.NewJtkHandler,
)

var WireSet = wire.NewSet(NewWireDefaultApp, handlerSet)
