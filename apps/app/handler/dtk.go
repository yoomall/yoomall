package handler

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core"
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

func (d *DtkHandler) Register(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Tag:    "dtk",
		Path:   "",
	}, d.dtk)
}

func (d *DtkHandler) GetRouterGroupName() string {
	return "dtk"
}

// 大淘客接口 godoc
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
