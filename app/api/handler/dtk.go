package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/app/model"
	"lazyfury.github.com/yoomall-server/constants"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/driver"
	"lazyfury.github.com/yoomall-server/libs/dtk"
)

type DtkHandler struct {
	*handler
	dtkClient *dtk.Dtk
}

func NewDtkHandler(app core.App) Handler {
	return &DtkHandler{
		handler: &handler{
			App: app,
		},
		dtkClient: dtk.NewDtkClient(app.GetConfig().DTK),
	}
}

func (d *DtkHandler) DB() *driver.DB {
	return constants.DB
}

func (d *DtkHandler) Register(router *gin.RouterGroup) {
	router.GET("/dtk-users", d.test)
	router.GET("/tb", d.dtk)
}

func (d *DtkHandler) test(ctx *gin.Context) {
	var users []model.User
	d.DB().Find(&users)
	ctx.JSON(http.StatusOK, map[string]any{"hello": users})
}

// 大淘客接口 godoc
//
//	@Summary		获取大淘客接口数据
//	@Description	大淘客接口
//	@Tags			/dtk/tb
//	@Accept			json
//	@Produce		json
//	@Param			keyWords	query	string	true	"keyWords"
//	@Param			pageSize	query	string	true	"pageSize"
//	@Param			pageId		query	string	true	"pageId"
//	@Router			/dtk/tb [get]
func (d *DtkHandler) dtk(ctx *gin.Context) {
	var query map[string]string = make(map[string]string)
	ctx.ShouldBindQuery(&query)
	log.Info("dtk", "query", query, "url", ctx.Request.URL.Query())
	resp, err := d.dtkClient.Request("/tb-service/get-tb-service", http.MethodGet, "v2.1.0", query)

	if err != nil {
		ctx.JSON(http.StatusOK, map[string]any{"error": err})
		return
	}

	body, err := io.ReadAll(resp.Body)
	print(string(body))
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]any{"error": err})
		return
	}

	defer resp.Body.Close()

	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]any{"hello": err})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{"hello": data})
}
