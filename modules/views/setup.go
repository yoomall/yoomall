package views

import (
	core "yoomall/yoo"
	"yoomall/yoo/driver"

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
	router.GET("", homeView)             // 首页
	router.GET("/about.html", aboutView) // 关于
}
