package handler

import (
	"net/http"

	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/apps/auth/model"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/curd"
)

type UserRoleHandler struct {
	DB        *driver.DB
	CRUD      *curd.CRUD
	authMidds *authmiddleware.AuthMiddlewareGroup
}

// GetRouterGroupName implements core.Handler.
func (u *UserRoleHandler) GetRouterGroupName() string {
	return "user-roles"
}

var _ core.Handler = (*UserRoleHandler)(nil)

func NewUserRoleHandler(db *driver.DB, authMidds *authmiddleware.AuthMiddlewareGroup) *UserRoleHandler {
	return &UserRoleHandler{
		DB:        db,
		CRUD:      curd.New(db, &model.UserRole{}),
		authMidds: authMidds,
	}
}

func (u *UserRoleHandler) Register(router *core.RouterGroup) {
	auth := router.Group("").Use(u.authMidds.MustAuthMiddleware)
	{
		auth.WithDoc(&core.DocItem{
			Method: http.MethodGet,
			Path:   "/role-list",
		}).GET("/role-list", u.CRUD.GetListHandler(&[]model.UserRole{}))
	}
}
