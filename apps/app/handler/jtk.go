package handler

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/helper/response"
	"lazyfury.github.com/yoomall-server/libs/jutuike"
)

type JtkHandler struct {
	client *jutuike.Jtk
}

func NewJtkHandler(config *viper.Viper) *JtkHandler {
	client, err := jutuike.NewJtkFromViper(config)
	if err != nil {
		log.Fatal(err)
	}
	return &JtkHandler{
		client: client,
	}
}

var _ core.Handler = (*JtkHandler)(nil)

func (j *JtkHandler) GetRouterGroupName() string {
	return "jtk"
}

func (j *JtkHandler) Register(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "",
	}, j.jtk)
}

// 聚推客接口 godoc
//
//	@Summary		获取聚推客接口数据
//	@Description	聚推客接口
//	@Tags			/jtk
//	@Accept			json
//	@Produce		json
//	@Param			path query string true "接口路径"
//	@Param			method query string true "请求方法"
//	@Param			...params query string false "请求参数/其他参数都是动态的参考聚推客开发文档/ swagger 不支持，请使用 apipost 工具调试"
//	@Router			/jtk [get]
func (j *JtkHandler) jtk(ctx *gin.Context) {
	var query map[string]string = make(map[string]string)
	ctx.ShouldBindQuery(&query)

	path := query["path"]
	if path == "" {
		response.Error(response.ErrBadRequest, "path is empty").Done(ctx)
		return
	}

	method := query["method"]
	if method == "" {
		method = http.MethodGet
	}

	resp, data, hit, err := j.client.RequestWithCache(path, method, "1", query)
	if err != nil {
		response.Error(response.ErrBadRequest, err.Error()).WithExtra(map[string]any{
			"response": data,
			"url":      resp.Request.URL.String(),
		}).Done(ctx)
		return
	}
	response.Success(data).WithExtra(map[string]any{
		"hit": hit,
	}).Done(ctx)
}
