package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"lazyfury.github.com/yoomall-server/app/handler"
	"lazyfury.github.com/yoomall-server/app/middleware"
	"lazyfury.github.com/yoomall-server/app/model"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

type DefaultApp struct {
	*core.AppImpl
	Config          *config.Config
	AuthMiddlewares []gin.HandlerFunc
}

var _ core.App = (*DefaultApp)(nil)

func NewWireDefaultApp(config *config.Config, db *driver.DB) *DefaultApp {
	return &DefaultApp{
		Config:          config,
		AppImpl:         core.NewAppImpl("default", config, db),
		AuthMiddlewares: []gin.HandlerFunc{},
	}
}

var WireSet = wire.NewSet(NewWireDefaultApp)

func (d *DefaultApp) Register(router *gin.RouterGroup) {
	handler.NewDtkHandler(d).Register(router.Group("/dtk"))
	handler.NewUserHandler(d).Register(router.Group("/users", d.AuthMiddlewares...))
}

func (d *DefaultApp) Migrate() {
	d.GetDB().AutoMigrate(&model.User{})
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CORSMiddleware(),
		coremiddleware.RecoverHandlerFunc,
	}
}
