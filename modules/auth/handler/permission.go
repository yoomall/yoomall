package handler

import (
	"net/http"

	authmiddleware "yoomall/modules/auth/middleware"
	"yoomall/modules/auth/model"
	"yoomall/yoo"
	"yoomall/yoo/driver"
	"yoomall/yoo/helper/curd"

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

func (p *PermissionHandler) Register(router *yoo.RouterGroup) {
	router.Use(p.authMidds.MustAuthMiddleware)
	router.WithDoc(&yoo.DocItem{
		Method: http.MethodGet,
		Path:   "/list",
	}).GET("/list", p.CRUD.GetListHandlerWithWhere(&[]model.Permission{}, func(tx *gorm.DB) *gorm.DB {
		return tx
	}))
}

func (p *PermissionHandler) GetRouterGroupName() string {
	return "permissions"
}
