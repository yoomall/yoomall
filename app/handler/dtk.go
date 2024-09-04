package handler

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/constants"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/response"
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
	router.GET("", d.dtk)
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
	body, hit := d.dtkClient.RequestWithCache("/tb-service/get-tb-service", http.MethodGet, "v2.1.0", query)

	var data map[string]any
	err := json.Unmarshal(body, &data)
	if err != nil {
		response.Error(response.ErrInternalError, err.Error()).Done(ctx)
		return
	}

	response.Success(data).WithExtra(map[string]any{
		"hit": hit,
	}).WithExtra(map[string]any{
		"hello": "world",
	}).Done(ctx)
}
