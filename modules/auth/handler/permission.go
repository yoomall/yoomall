package handler

import (
	"net/http"

	authmiddleware "yoomall/modules/auth/middleware"
	"yoomall/modules/auth/model"

	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/lazyfury/pulse/helper/curd"
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

func (p *PermissionHandler) Register(router *framework.RouterGroup) {
	router.Use(p.authMidds.MustAuthMiddleware)
	router.GET("/list", p.CRUD.GetListHandlerWithWhere(&[]model.Permission{}, func(tx *gorm.DB) *gorm.DB {
		return tx
	})).Doc(&framework.DocItem{
		Method: http.MethodGet,
		Path:   "/list",
	})
}

func (p *PermissionHandler) GetRouterGroupName() string {
	return "permissions"
}
