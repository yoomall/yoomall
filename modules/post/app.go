package post

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
)

type DefaultApp struct {
	*core.AppImpl
	Config *config.Config
}

func NewDefaultApp(engine *gin.Engine, router *gin.RouterGroup, config *config.Config) core.App {
	return &DefaultApp{
		Config:  config,
		AppImpl: core.NewAppImpl("post", router, config),
	}
}

func (d *DefaultApp) Register() {
	d.GetRouter().GET("/list", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{"data": []any{}})
	})
}

func (d *DefaultApp) Migrate() {
	d.GetDB().AutoMigrate()
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
