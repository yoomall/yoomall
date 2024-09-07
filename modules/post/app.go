package post

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type DefaultApp struct {
	*core.AppImpl
	Config *config.Config
}

func NewDefaultApp(config *config.Config, db *driver.DB) *DefaultApp {
	return &DefaultApp{
		Config:  config,
		AppImpl: core.NewAppImpl("post", config, db, []core.Handler{}),
	}
}

var WireSet = wire.NewSet(NewDefaultApp)
var _ core.App = (*DefaultApp)(nil)

func (d *DefaultApp) Register(router *gin.RouterGroup) {
	router.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{"data": []any{}})
	})
}

func (d *DefaultApp) Migrate() {
	d.GetDB().AutoMigrate()
}

func (d *DefaultApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
