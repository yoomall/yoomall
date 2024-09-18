package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type PostApp struct {
	*core.AppImpl
	Config *viper.Viper
}

func NewDefaultApp(config *viper.Viper, db *driver.DB) *PostApp {
	return &PostApp{
		Config:  config,
		AppImpl: core.NewAppImpl("post", config, db, []core.Handler{}),
	}
}

var WireSet = wire.NewSet(NewDefaultApp)
var _ core.App = (*PostApp)(nil)

func (d *PostApp) Register(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Tag:    "post",
		Path:   "/list",
	}, func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{"data": []any{}})
	})
}

func (d *PostApp) Migrate() {
	d.GetDB().AutoMigrate()
}

func (d *PostApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
