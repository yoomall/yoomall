package core

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/constants"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type App interface {
	GetName() string
	GetRouter() *gin.RouterGroup
	Register() // 注册路由
	GetDB() *driver.DB
	Migrate()
	Middleware() []gin.HandlerFunc
	GetConfig() *config.Config
}

type AppImpl struct {
	AppName   string
	AppRouter *gin.RouterGroup
	AppConfig *config.Config
}

func NewAppImpl(name string, router *gin.RouterGroup, config *config.Config) *AppImpl {
	return &AppImpl{
		AppName:   name,
		AppRouter: router,
		AppConfig: config,
	}
}

func (a *AppImpl) GetConfig() *config.Config {
	return a.AppConfig
}

func (a *AppImpl) GetName() string {
	return a.AppName
}

func (a *AppImpl) GetDB() *driver.DB {
	return constants.DB
}

func (a *AppImpl) GetRouter() *gin.RouterGroup {
	return a.AppRouter
}
