package handler

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core/helper/response"
	"lazyfury.github.com/yoomall-server/libs/dtk"
)

type DtkHandler struct {
	client *dtk.Dtk
}

func NewDtkHandler(config *viper.Viper) *DtkHandler {
	clent, err := dtk.NewDtkFromViper(config)
	if err != nil {
		log.Fatal(err)
	}
	return &DtkHandler{
		client: clent,
	}
}

func (d *DtkHandler) Register(router *gin.RouterGroup) {
	router.GET("", d.dtk)
}

func (d *DtkHandler) GetRouterGroupName() string {
	return "dtk"
}

// 大淘客接口 godoc
//
//		@Summary		获取大淘客接口数据
//		@Description	大淘客接口
//		@Tags			/dtk
//		@Accept			json
//		@Produce		json
//		@Param			path query string true "接口路径"
//		@Param			method query string true "请求方法"
//		@Param			...params query string false "请求参数/其他参数都是动态的参考聚推客开发文档/ swagger 不支持，请使用 apipost 工具调试"
//		@Router			/dtk [get]
//	 @Success		200 {object} response.ApiJsonResponse
//	 @Failure		500 {object} response.ApiJsonResponse
func (d *DtkHandler) dtk(ctx *gin.Context) {
	var query map[string]string = make(map[string]string)
	ctx.ShouldBindQuery(&query)

	path := query["path"]
	version := query["version"]
	if path == "" || version == "" {
		response.Error(response.ErrBadRequest, "path or version is empty").Done(ctx)
		return
	}

	method := query["method"]
	if method == "" {
		method = http.MethodGet
	}

	delete(query, "method")
	delete(query, "path")
	delete(query, "version")

	log.Info("dtk", "query", query, "url", ctx.Request.URL.Query())
	resp, data, hit, err := d.client.RequestWithCache(path, method, version, query)

	extra := map[string]any{
		"hit": hit,
	}

	if err != nil {
		response.Error(response.ErrInternalError, err.Error()).WithExtra(extra).WithExtra(map[string]any{
			"text":   data,
			"url":    resp.Request.URL.String(),
			"method": method,
		}).Done(ctx)
		return
	}

	response.Success(data).WithExtra(extra).WithExtra(map[string]any{}).Done(ctx)
}
