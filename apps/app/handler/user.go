package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/apps/app/model"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/curd"
)

type UserHandler struct {
	CRUD *curd.CRUD
}

var _ core.Handler = (*UserHandler)(nil)

func NewUserHandler(db *driver.DB, config *viper.Viper) *UserHandler {
	return &UserHandler{
		CRUD: &curd.CRUD{
			DB:    db,
			Model: &model.User{},
		},
	}
}

func (u *UserHandler) Register(router *gin.RouterGroup) {
	router.GET("/list", u.CRUD.GetListHandler(&[]model.User{}))
}

func (u *UserHandler) GetRouterGroupName() string {
	return "users"
}
