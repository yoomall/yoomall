package ui

import (
	"encoding/json"

	"lazyfury.github.com/yoomall-server/core"
)

type IRouter interface {
	GetPath() string
	GetParams() map[string]string
	GetComponentName() string
}

type Router struct {
	Key    string            `json:"key"`
	Title  string            `json:"title"`
	Icon   string            `json:"icon"`
	Path   string            `json:"path"`
	Params map[string]string `json:"params"`
	Page   IPage             `json:"page"`
}

var _ IRouter = (*Router)(nil)
var _ core.MarshalJSON = (*Router)(nil)

func (r *Router) GetPath() string {
	return r.Path
}

func (r *Router) GetParams() map[string]string {
	return r.Params
}

func (r *Router) GetComponentName() string {
	return r.Page.GetComponentName()
}

func (r *Router) MarshalJSON() ([]byte, error) {
	type Alias Router
	type A struct {
		*Alias
		Component string `json:"component"`
	}
	var a *A = &A{
		Alias:     (*Alias)(r),
		Component: r.Page.GetComponentName(),
	}
	return json.Marshal(a)
}

func NewRouter(key, title, icon, path string, params map[string]string, page IPage) *Router {
	return &Router{
		Key:    key,
		Title:  title,
		Icon:   icon,
		Path:   path,
		Params: params,
		Page:   page,
	}
}
