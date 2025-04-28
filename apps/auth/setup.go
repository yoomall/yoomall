package auth

import (
	"github.com/yoomall/yoomall/apps/auth/handler"
	authmiddleware "github.com/yoomall/yoomall/apps/auth/middleware"
	"github.com/yoomall/yoomall/apps/auth/model"
	authservice "github.com/yoomall/yoomall/apps/auth/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/spf13/viper"
)

type AuthApp struct {
	*framework.App

	service *authservice.AuthService
}

var _ framework.IApp = (*AuthApp)(nil)

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
		App: framework.NewApp("auth", config, db, []framework.Handler{
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

func (a *AuthApp) Register(router *framework.RouterGroup) {

}
