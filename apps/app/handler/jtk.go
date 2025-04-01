package handler

import (
	"net/http"

	"yoomall/libs/jutuike"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/helper/response"
	"github.com/spf13/viper"
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

var _ framework.Handler = (*JtkHandler)(nil)

func (j *JtkHandler) GetRouterGroupName() string {
	return "jtk"
}

func (j *JtkHandler) Register(router *framework.RouterGroup) {
	router.GET("", j.jtk).Doc(&framework.DocItem{
		Method: http.MethodGet,
		Path:   "",
	})
}

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
