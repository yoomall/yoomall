package core

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type App interface {
	GetName() string
	GetDB() *driver.DB
	Migrate()
	Middleware() []gin.HandlerFunc
	GetConfig() *viper.Viper
	Register(router *gin.RouterGroup)
}

type AppImpl struct {
	AppName   string
	AppConfig *viper.Viper
	db        *driver.DB
	Handlers  []Handler
}

func NewAppImpl(name string, config *viper.Viper, db *driver.DB, handlers []Handler) *AppImpl {
	return &AppImpl{
		AppName:   name,
		AppConfig: config,
		db:        db,
		Handlers:  handlers,
	}
}

func (a *AppImpl) GetConfig() *viper.Viper {
	return a.AppConfig
}

func (a *AppImpl) GetName() string {
	return a.AppName
}

func (a *AppImpl) GetDB() *driver.DB {
	return a.db
}
