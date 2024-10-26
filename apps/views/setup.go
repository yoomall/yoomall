package views

import (
	"yoomall/apps/views/render"
	"yoomall/core"
	"yoomall/core/driver"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type ViewsApp struct {
	*core.App
}

var _ core.IApp = (*ViewsApp)(nil)

func NewViewApp(db *driver.DB, config *viper.Viper) *ViewsApp {
	return &ViewsApp{
		App: core.NewApp("view", config, db, []core.Handler{}).WithPlugins(nil),
	}
}

var WireSet = wire.NewSet(NewViewApp)

// Migrate implements core.IApp.
func (v *ViewsApp) Migrate() {
}

// Register implements core.IApp.
func (v *ViewsApp) Register(router *core.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {
		render.Html(ctx, "index.html", nil)
	})

	router.GET("/about.html", func(ctx *gin.Context) {
		render.Html(ctx, "about.html", nil)
	})
}
