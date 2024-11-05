package handler

import (
	"net/http"

	"yoomall/libs/dtk"
	"yoomall/yoo"
	"yoomall/yoo/helper/response"
	"yoomall/yoo/helper/utils"
	"yoomall/yoo/helper/validate"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func (d *DtkHandler) Register(router *yoo.RouterGroup) {
	router.WithDoc(&yoo.DocItem{
		Method: http.MethodGet,
		Tag:    "dtk",
		Path:   "",
	}).GET("", d.dtk)
}

func (d *DtkHandler) GetRouterGroupName() string {
	return "dtk"
}

// 大淘客接口 godoc
func (d *DtkHandler) dtk(ctx *gin.Context) {
	var query map[string]string = make(map[string]string)
	ctx.ShouldBindQuery(&query)

	validator := validate.NewValidator()
	validator.AddValidate(validate.NewStringValidate("path", false, "path is empty", 2, 300, nil))
	validator.AddValidate(validate.NewStringValidate("version", false, "version is empty", 0, 0, nil))
	if valid, msg := validator.Validate(utils.StringMapToInterfaceMap(query)); !valid {
		response.Error(response.ErrBadRequest, msg).Done(ctx)
		return
	}

	path := query["path"]
	version := query["version"]
	method := utils.GetFromMapWithDefault(query, "method", http.MethodGet)

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
