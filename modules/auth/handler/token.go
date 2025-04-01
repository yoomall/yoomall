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

var _ framework.Handler = (*UserTokenHandler)(nil)

func (u *UserTokenHandler) Register(router *framework.RouterGroup) {
	router.Use(u.authMidds.MustAuthMiddleware)
	router.GET("/logs", u.CRUD.GetListHandlerWithWhere(&[]model.UserToken{}, func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("User")
	})).Doc(&framework.DocItem{
		Method: http.MethodGet,
		Path:   "/logs",
	})
}
