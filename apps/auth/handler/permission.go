package handler

import (
	"net/http"

	"gorm.io/gorm"
	authmiddleware "lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/apps/auth/model"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/curd"
)

type PermissionHandler struct {
	DB        *driver.DB
	authMidds *authmiddleware.AuthMiddlewareGroup
	CRUD      *curd.CRUD
}

func NewPermissionHandler(db *driver.DB, authMidds *authmiddleware.AuthMiddlewareGroup) *PermissionHandler {
	return &PermissionHandler{
		DB:        db,
		authMidds: authMidds,
		CRUD:      curd.New(db, &model.Permission{}),
	}
}

func (p *PermissionHandler) Register(router *core.RouterGroup) {
	router.Use(p.authMidds.MustAuthMiddleware)
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "/list",
	}).GET("/list", p.CRUD.GetListHandlerWithWhere(&[]model.Permission{}, func(tx *gorm.DB) *gorm.DB {
		return tx
	}))
}

func (p *PermissionHandler) GetRouterGroupName() string {
	return "permissions"
}
