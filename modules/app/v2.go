package app

import (
	appHandlerV2 "yoomall/modules/app/handler/v2"
	"yoomall/yoo"
	"yoomall/yoo/driver"

	"github.com/spf13/viper"
)

type DefaultV2App struct {
	*yoo.App
}

var _ yoo.IApp = (*DefaultV2App)(nil)

func NewDefaultV2App(config *viper.Viper, db *driver.DB,
	baseV2Handler *appHandlerV2.BaseHandlerV2,
) *DefaultV2App {
	return &DefaultV2App{
		App: yoo.NewApp("defaultV2", config, db, []yoo.Handler{
			baseV2Handler,
		}),
	}
}

// Migrate implements yoo.IApp.
func (d *DefaultV2App) Migrate() {

}

// Register implements yoo.IApp.
func (d *DefaultV2App) Register(router *yoo.RouterGroup) {

}
