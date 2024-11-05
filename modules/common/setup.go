package common

import (
	"yoomall/modules/common/handler"
	"yoomall/modules/common/model"
	commonservice "yoomall/modules/common/service"
	yoo "yoomall/yoo"
	"yoomall/yoo/driver"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type CommonApp struct {
	*yoo.App
}

var _ yoo.IApp = (*CommonApp)(nil)

// GetName implements yoo.App.
// Subtle: this method shadows the method (*AppImpl).GetName of CommonApp.AppImpl.
func (c *CommonApp) GetName() string {
	return "common"
}

// Middleware implements yoo.App.
func (c *CommonApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

// Migrate implements yoo.App.
func (c *CommonApp) Migrate() {
	c.GetDB().AutoMigrate(
		&model.NotFoundRecord{},
		&model.SystemConfig{},
		&model.SystemConfigGroup{},
	)
}

// Register implements yoo.App.
func (c *CommonApp) Register(router *yoo.RouterGroup) {

}

func NewCommonApp(config *viper.Viper, db *driver.DB,
	notfoundHandler *handler.NotFoundRecordHandler,
	systemConfigHandler *handler.SystemConfigHandler,
) *CommonApp {
	return &CommonApp{
		App: yoo.NewApp("common", config, db, []yoo.Handler{
			notfoundHandler,
			systemConfigHandler,
		}),
	}
}

var WireSet = wire.NewSet(NewCommonApp, commonservice.NewNotFoundRecordService, handler.NewNotFoundRecordHandler, handler.NewSystemConfigHandler, commonservice.NewSystemConfigService)
