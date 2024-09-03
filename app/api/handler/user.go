package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/app/model"
	"lazyfury.github.com/yoomall-server/core"
)

type UserHandler struct {
	*handler
}

func NewUserHandler(app core.App) Handler {
	return &UserHandler{
		handler: &handler{
			App: app,
		},
	}
}

func (u *UserHandler) Register(router *gin.RouterGroup) {
	router.GET("/users", u.users)
}

func (u *UserHandler) users(ctx *gin.Context) {
	var users []model.User
	u.App.GetDB().Find(&users)
	ctx.JSON(http.StatusOK, map[string]any{"users": users})
}
