package common

import (
	"github.com/yoomall/yoomall/apps/common/handler"
	"github.com/yoomall/yoomall/apps/common/model"
	commonservice "github.com/yoomall/yoomall/apps/common/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

type CommonApp struct {
	*framework.App
}

var _ framework.IApp = (*CommonApp)(nil)

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
func (c *CommonApp) Register(router *framework.RouterGroup) {

}

func NewCommonApp(config *viper.Viper, db *driver.DB,
	notfoundHandler *handler.NotFoundRecordHandler,
	systemConfigHandler *handler.SystemConfigHandler,
) *CommonApp {
	return &CommonApp{
		App: framework.NewApp("common", config, db, []framework.Handler{
			notfoundHandler,
			systemConfigHandler,
		}),
	}
}

var WireSet = wire.NewSet(NewCommonApp, commonservice.NewNotFoundRecordService, handler.NewNotFoundRecordHandler, handler.NewSystemConfigHandler, commonservice.NewSystemConfigService)
