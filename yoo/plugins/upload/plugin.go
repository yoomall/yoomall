package upload

import (
	"net/http"
	"yoomall/yoo"
	"yoomall/yoo/helper/response"

	"github.com/gin-gonic/gin"
)

type UploadPlugin struct {
	*yoo.Plugin
	uploader *Uploader
}

var _ yoo.IPlugin = (*UploadPlugin)(nil)

func NewUploadPlugin() *UploadPlugin {
	return &UploadPlugin{
		Plugin: yoo.NewPlugin(),
		uploader: &Uploader{
			BaseDir:      "./static/upload",
			UploadMethod: DefaultUpload,
			GetFile:      DefaultGetFile,
		},
	}
}

func (p *UploadPlugin) RegisterRouter(router *yoo.RouterGroup) {
	router.WithDoc(&yoo.DocItem{
		Method: http.MethodPost,
		Path:   "/upload",
	}).POST("/upload", p.upload)
}

func (p *UploadPlugin) upload(ctx *gin.Context) {
	path, err := p.uploader.Default(ctx.Request)
	if err != nil {
		response.Error(response.ErrBadRequest, err.Error()).Done(ctx)
		return
	}
	response.Success(map[string]any{
		"url": path,
	}).Done(ctx)
}
