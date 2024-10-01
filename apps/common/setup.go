package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/common/handler"
	"lazyfury.github.com/yoomall-server/apps/common/model"
	commonservice "lazyfury.github.com/yoomall-server/apps/common/service"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type CommonApp struct {
	*core.App
}

var _ core.IApp = (*CommonApp)(nil)

// GetName implements core.App.
// Subtle: this method shadows the method (*AppImpl).GetName of CommonApp.AppImpl.
func (c *CommonApp) GetName() string {
	return "common"
}

// Middleware implements core.App.
func (c *CommonApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

// Migrate implements core.App.
func (c *CommonApp) Migrate() {
	c.GetDB().AutoMigrate(
		&model.NotFoundRecord{},
		&model.SystemConfig{},
		&model.SystemConfigGroup{},
	)
}

// Register implements core.App.
func (c *CommonApp) Register(router *core.RouterGroup) {

}

func NewCommonApp(config *viper.Viper, db *driver.DB,
	notfoundHandler *handler.NotFoundRecordHandler,
	systemConfigHandler *handler.SystemConfigHandler,
) *CommonApp {
	return &CommonApp{
		App: core.NewApp("common", config, db, []core.Handler{
			notfoundHandler,
			systemConfigHandler,
		}),
	}
}

var WireSet = wire.NewSet(NewCommonApp, commonservice.NewNotFoundRecordService, handler.NewNotFoundRecordHandler, handler.NewSystemConfigHandler, commonservice.NewSystemConfigService)
