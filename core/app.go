package core

import (
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/constants"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type App interface {
	GetName() string
	Register(router *gin.RouterGroup) // 注册路由
	GetDB() *driver.DB
	Migrate()
	Middleware() []gin.HandlerFunc
	GetConfig() *config.Config
}

type AppImpl struct {
	AppName   string
	AppConfig *config.Config
	db        *driver.DB
}

func NewAppImpl(name string, config *config.Config, db *driver.DB) *AppImpl {
	return &AppImpl{
		AppName:   name,
		AppConfig: config,
		db:        db,
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
