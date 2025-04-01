package app

import (
	appHandlerV2 "yoomall/apps/app/handler/v2"

	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

type DefaultV2App struct {
	*framework.App
}

var _ framework.IApp = (*DefaultV2App)(nil)

func NewDefaultV2App(config *viper.Viper, db *driver.DB,
	baseV2Handler *appHandlerV2.BaseHandlerV2,
) *DefaultV2App {
	return &DefaultV2App{
		App: framework.NewApp("defaultV2", config, db, []framework.Handler{
			baseV2Handler,
		}),
	}
}

// Migrate implements yoo.IApp.
func (d *DefaultV2App) Migrate() {

}

// Register implements yoo.IApp.
func (d *DefaultV2App) Register(router *framework.RouterGroup) {

}
