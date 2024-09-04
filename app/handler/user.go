package handler

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/app/model"
	"lazyfury.github.com/yoomall-server/core"
)

type UserHandler struct {
	*handler
	CRUD *core.CRUD
}

func NewUserHandler(app core.App) Handler {
	return &UserHandler{
		handler: &handler{
			App: app,
		},
		CRUD: &core.CRUD{
			DB:    app.GetDB(),
			Model: &model.User{},
		},
	}
}

func (u *UserHandler) Register(router *gin.RouterGroup) {
	router.GET("", u.CRUD.GetListHandler(&[]model.User{}))
}
