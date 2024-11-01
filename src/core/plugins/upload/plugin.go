package upload

import (
	"net/http"
	"yoomall/src/core"
	"yoomall/src/core/helper/response"

	"github.com/gin-gonic/gin"
)

type UploadPlugin struct {
	*core.Plugin
	uploader *Uploader
}

var _ core.IPlugin = (*UploadPlugin)(nil)

func NewUploadPlugin() *UploadPlugin {
	return &UploadPlugin{
		Plugin: core.NewPlugin(),
		uploader: &Uploader{
			BaseDir:      "./static/upload",
			UploadMethod: DefaultUpload,
			GetFile:      DefaultGetFile,
		},
	}
}

func (p *UploadPlugin) RegisterRouter(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
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
