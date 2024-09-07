package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/app/handler"
	"lazyfury.github.com/yoomall-server/apps/app/service"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

func NewWireDefaultApp(config *viper.Viper, db *driver.DB,
	userHandler *handler.UserHandler,
	dtkHandler *handler.DtkHandler,
	menuHandler *handler.MenuHandler,
) *DefaultApp {
	return &DefaultApp{
		Config: config,
		AppImpl: core.NewAppImpl("default", config, db, []core.Handler{
			userHandler,
			dtkHandler,
			menuHandler,
		}),
		AuthMiddlewares: []gin.HandlerFunc{},
	}
}

var handlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewDtkHandler,
	handler.NewMenuHandler,
)
var serviceSet = wire.NewSet(service.NewAuthService)

var WireSet = wire.NewSet(NewWireDefaultApp, handlerSet, serviceSet)
