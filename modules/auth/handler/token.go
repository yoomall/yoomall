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

type UserTokenHandler struct {
	DB        *driver.DB
	authMidds *authmiddleware.AuthMiddlewareGroup
	CRUD      *curd.CRUD
}

// GetRouterGroupName implements yoo.Handler.
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

var _ yoo.Handler = (*UserTokenHandler)(nil)

func (u *UserTokenHandler) Register(router *yoo.RouterGroup) {
	router.Use(u.authMidds.MustAuthMiddleware)
	router.WithDoc(&yoo.DocItem{
		Method: http.MethodGet,
		Path:   "/logs",
	}).GET("/logs", u.CRUD.GetListHandlerWithWhere(&[]model.UserToken{}, func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("User")
	}))
}
