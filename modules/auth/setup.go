package auth

import (
	"yoomall/modules/auth/handler"
	authmiddleware "yoomall/modules/auth/middleware"
	"yoomall/modules/auth/model"
	authservice "yoomall/modules/auth/service"
	core "yoomall/yoo"
	"yoomall/yoo/driver"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type AuthApp struct {
	*core.App

	service *authservice.AuthService
}

var _ core.IApp = (*AuthApp)(nil)

var handlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewUserRoleHandler,
	handler.NewUserTokenHandler,
	handler.NewPermissionHandler,
)
var WireSet = wire.NewSet(
	NewAuthApp, authservice.NewAuthService, authmiddleware.NewAuthMiddlewareGroup,
	handlerSet,
)

func NewAuthApp(
	config *viper.Viper,
	db *driver.DB,
	service *authservice.AuthService,
	userHandler *handler.UserHandler,
	roleHandler *handler.UserRoleHandler,
	tokenHandler *handler.UserTokenHandler,
	permissionHandler *handler.PermissionHandler,
) *AuthApp {
	return &AuthApp{
		App: core.NewApp("auth", config, db, []core.Handler{
			userHandler,
			roleHandler,
			tokenHandler,
			permissionHandler,
		}),

		service: service,
	}
}

func (a *AuthApp) Migrate() {
	a.GetDB().AutoMigrate(
		&model.User{}, &model.UserToken{},
		&model.UserRole{}, &model.UserRoleRef{},
		&model.Permission{}, &model.RolePermissionRef{}, &model.UserPermissionRef{},
	)
}

func (a *AuthApp) Middleware() []gin.HandlerFunc {
	return nil
}

func (a *AuthApp) Register(router *core.RouterGroup) {

}
