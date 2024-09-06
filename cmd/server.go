package main

import (
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lazyfury.github.com/yoomall-server/app"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/constants"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/docs"
	"lazyfury.github.com/yoomall-server/modules/post"
)

// @title						Nunu Example API
// @version					1.0.0
// @description				This is a sample server celler server.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8900
// @securityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	log.Info("hello world")
	instance := gin.Default()
	instance.SetTrustedProxies(nil)

	// config
	constants.CONFIG = config.GetConfig("./config.yaml")
	log.Info("config load success", "config", constants.CONFIG)

	if constants.CONFIG.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// db
	constants.DB = driver.NewDB(constants.CONFIG.MysqlDsn())
	// db

	docs.SwaggerInfo.BasePath = "/api"
	instance.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true)))

	router := instance.Group("/v1/api")
	// apps
	defApp := app.NewDefaultApp(instance, router.Group("/"), constants.CONFIG)
	postApp := post.NewDefaultApp(instance, router.Group("/posts"), constants.CONFIG)
	register(defApp, postApp)
	port := strconv.Itoa(constants.CONFIG.HTTP.Port)
	if port == "0" {
		port = "8900"
	}
	log.Info("[server on] http://127.0.0.1:" + port)
	instance.Run(":" + port)
}

func register(apps ...core.App) {

	for _, app := range apps {
		log.Info("register app", "app", app.GetName())
		if constants.CONFIG.DEBUG {
			app.Migrate()
			log.Info("migrate success", "app", app.GetName())
		}
		app.GetRouter().Use(app.Middleware()...)
		app.Register()
		log.Info("register app success", "app", app.GetName())
	}
}
