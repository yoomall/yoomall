package views

import (
	"github.com/google/wire"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

type ViewsApp struct {
	*framework.App
}

var _ framework.IApp = (*ViewsApp)(nil)

func NewViewApp(db *driver.DB, config *viper.Viper) *ViewsApp {
	return &ViewsApp{
		App: framework.NewApp("view", config, db, []framework.Handler{}).WithPlugins(nil),
	}
}

var WireSet = wire.NewSet(NewViewApp)

// Migrate implements core.IApp.
func (v *ViewsApp) Migrate() {
}

// Register implements core.IApp.
func (v *ViewsApp) Register(router *framework.RouterGroup) {
	router.GET("", homeView)             // 首页
	router.GET("/about.html", aboutView) // 关于
}
