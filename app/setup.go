package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	defappConfig "lazyfury.github.com/yoomall-server/app/config"
	"lazyfury.github.com/yoomall-server/app/handler"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

func NewWireDefaultApp(config *config.Config, db *driver.DB,
	userHandler *handler.UserHandler,
	dtkHandler *handler.DtkHandler,
) *defappConfig.DefaultApp {
	return &defappConfig.DefaultApp{
		Config: config,
		AppImpl: core.NewAppImpl("default", config, db, []core.Handler{
			userHandler,
			dtkHandler,
		}),
		AuthMiddlewares: []gin.HandlerFunc{},
	}
}

var WireSet = wire.NewSet(NewWireDefaultApp, handler.NewUserHandler, handler.NewDtkHandler)
