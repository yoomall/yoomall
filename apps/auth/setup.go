package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/auth/handler"
	"lazyfury.github.com/yoomall-server/apps/auth/model"
	"lazyfury.github.com/yoomall-server/apps/auth/service"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type AuthApp struct {
	*core.AppImpl
}

var _ core.App = (*AuthApp)(nil)

var WireSet = wire.NewSet(NewAuthApp, service.NewAuthService, handler.NewUserHandler)

func NewAuthApp(
	config *viper.Viper,
	db *driver.DB,
	userHandler *handler.UserHandler,

) *AuthApp {
	return &AuthApp{
		AppImpl: core.NewAppImpl("auth", config, db, []core.Handler{
			userHandler,
		}),
	}
}

func (a *AuthApp) Migrate() {
	a.GetDB().AutoMigrate(&model.User{}, &model.UserToken{}, &model.UserExt{})
}

func (a *AuthApp) Middleware() []gin.HandlerFunc {
	return nil
}

func (a *AuthApp) Register(router *core.RouterGroup) {

}
