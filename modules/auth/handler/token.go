package handler

import (
	"net/http"

	"yoomall/core"
	"yoomall/core/driver"
	"yoomall/core/helper/curd"
	authmiddleware "yoomall/modules/auth/middleware"
	"yoomall/modules/auth/model"

	"gorm.io/gorm"
)

type UserTokenHandler struct {
	DB        *driver.DB
	authMidds *authmiddleware.AuthMiddlewareGroup
	CRUD      *curd.CRUD
}

// GetRouterGroupName implements core.Handler.
func (u *UserTokenHandler) GetRouterGroupName() string {
	return "user-logs"
}

func NewUserTokenHandler(db *driver.DB, authMidds *authmiddleware.AuthMiddlewareGroup) *UserTokenHandler {
	return &UserTokenHandler{
		DB:        db,
		authMidds: authMidds,
		CRUD:      curd.New(db, &model.UserToken{}),
	}
}

var _ core.Handler = (*UserTokenHandler)(nil)

func (u *UserTokenHandler) Register(router *core.RouterGroup) {
	router.Use(u.authMidds.MustAuthMiddleware)
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "/logs",
	}).GET("/logs", u.CRUD.GetListHandlerWithWhere(&[]model.UserToken{}, func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("User")
	}))
}
