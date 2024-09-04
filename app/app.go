package app

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/app/handler"
	"lazyfury.github.com/yoomall-server/app/middleware"
	"lazyfury.github.com/yoomall-server/app/model"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	coremiddleware "lazyfury.github.com/yoomall-server/core/middleware"
)

type DefaultApp struct {
	*core.AppImpl
	Config *config.Config
}

func NewDefaultApp(engine *gin.Engine, router *gin.RouterGroup, config *config.Config) core.App {
	return &DefaultApp{
		Config:  config,
		AppImpl: core.NewAppImpl("default", router, config),
	}
}

func (d *DefaultApp) Register() {
	handler.NewDtkHandler(d).Register(d.GetRouter().Group("/dtk"))
	handler.NewUserHandler(d).Register(d.GetRouter().Group("/users"))
}

func (d *DefaultApp) Migrate() {
	d.GetDB().AutoMigrate(&model.User{})
	log.Info("default app", "db", "auto migrate success")
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CORSMiddleware(),
		coremiddleware.RecoverHandlerFunc,
	}
}
