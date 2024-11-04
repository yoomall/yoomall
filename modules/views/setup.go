package views

import (
	"net/http"
	"yoomall/core"
	"yoomall/core/driver"
	"yoomall/modules/views/render"

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
		render.Html("index.html", map[string]any{
			"packages": []any{
				map[string]any{
					"price": 0,
				},
				map[string]any{
					"price": 680,
				},
				map[string]any{
					"price": 1980,
				},
			},
			"products": []any{
				map[string]any{
					"name": "yoomall",
				},
				map[string]any{
					"name": "yoomall-ui",
				},
				map[string]any{
					"name": "yoomall-admin",
				},
				map[string]any{
					"name": "yoomall-admin-ui",
				},
				map[string]any{
					"name": "yoomall-admin-api",
				},
				map[string]any{
					"name": "yoomall-admin-console",
				},
				map[string]any{
					"name": "yoomall-admin-ui",
				},
			},
		}).SEO("easyly make your design to production.", "yoomall", "yoomall").Done(ctx)
	})

	router.GET("/about.html", func(ctx *gin.Context) {
		render.Html("about.html", nil).Done(ctx)
	})
}

func (v *ViewsApp) NotFoundHandler(ctx *gin.Context) {
	if ctx.Request.Header.Get("Accept") == "application/json" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "不存在的路由"})
		return
	}
	render.Html("404.html", nil).Done(ctx)
}
