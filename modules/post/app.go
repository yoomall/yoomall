package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

type PostApp struct {
	*framework.App
	Config *viper.Viper
}

func NewDefaultApp(config *viper.Viper, db *driver.DB) *PostApp {
	return &PostApp{
		Config: config,
		App:    framework.NewApp("post", config, db, []framework.Handler{}),
	}
}

var WireSet = wire.NewSet(NewDefaultApp)
var _ framework.IApp = (*PostApp)(nil)

func (d *PostApp) Register(router *framework.RouterGroup) {
	router.Doc(&framework.DocItem{
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
