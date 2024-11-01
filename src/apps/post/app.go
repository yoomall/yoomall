package post

import (
	"net/http"

	"yoomall/src/core"
	"yoomall/src/core/driver"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type PostApp struct {
	*core.App
	Config *viper.Viper
}

func NewDefaultApp(config *viper.Viper, db *driver.DB) *PostApp {
	return &PostApp{
		Config: config,
		App:    core.NewApp("post", config, db, []core.Handler{}),
	}
}

var WireSet = wire.NewSet(NewDefaultApp)
var _ core.IApp = (*PostApp)(nil)

func (d *PostApp) Register(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Tag:    "post",
		Path:   "/list",
	}).GET("/list", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{"data": []any{}})
	})
}

func (d *PostApp) Migrate() {
	d.GetDB().AutoMigrate()
}

func (d *PostApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
