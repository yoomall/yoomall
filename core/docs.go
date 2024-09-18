package core

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

type Doc struct {
	Title   string
	URL     string
	Version string
	Items   []*DocItem
}

// MarshalJSON implements MarshalJSON.
// 可能需要整理成 openapi 适合的结构，暂时不处理
func (d *Doc) MarshalJSON() ([]byte, error) {
	type Alias Doc
	var a = &struct {
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	return json.Marshal(a)
}

var _ MarshalJSON = (*Doc)(nil)

type DocItem struct {
	Title       string
	Content     string
	Description string
	Tag         string
	Path        string
	Method      string
	Params      any
	Body        any
	Success     any
	Failure     any
}

var DocInstance *Doc

func NewDoc() *Doc {
	DocInstance = &Doc{}
	return DocInstance
}

func (d *Doc) Add(item *DocItem) {
	d.Items = append(d.Items, item)
}

func (d *Doc) Handler(ctx *gin.Context) {
	response.Success(d).Done(ctx)
}
