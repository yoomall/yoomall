package handler

import (
	"net/http"

	authmiddleware "yoomall/src/apps/auth/middleware"
	"yoomall/src/apps/auth/model"
	"yoomall/src/core"
	"yoomall/src/core/driver"
	"yoomall/src/core/helper/curd"

	"gorm.io/gorm"
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
