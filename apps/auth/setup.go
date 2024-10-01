package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/auth/handler"
	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/apps/auth/model"
	"lazyfury.github.com/yoomall-server/apps/auth/service"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

type AuthApp struct {
	*core.App

	service *service.AuthService
}

var _ core.IApp = (*AuthApp)(nil)

var handlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewUserRoleHandler,
	handler.NewUserTokenHandler,
	handler.NewPermissionHandler,
)
var WireSet = wire.NewSet(
	NewAuthApp, service.NewAuthService, authmiddleware.NewAuthMiddlewareGroup,
	handlerSet,
)

func NewAuthApp(
	config *viper.Viper,
	db *driver.DB,
	service *service.AuthService,
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
	auth := router.Group("").Use(authmiddleware.NewMustAuthMiddlewareWithUser(a.GetDB()))
	{
		auth.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/profile",
		}).GET("/profile", func(ctx *gin.Context) {
			response.Success(ctx.MustGet("user")).Done(ctx)
		})

		auth.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/logout",
		}).POST("/logout", func(ctx *gin.Context) {
			a.service.Logout(ctx)
		})
	}
}
