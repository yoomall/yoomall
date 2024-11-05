package handler

import (
	"net/http"

	"yoomall/modules/common/model"
	commonservice "yoomall/modules/common/service"
	"yoomall/yoo"
	"yoomall/yoo/driver"
	"yoomall/yoo/helper/curd"
)

type NotFoundRecordHandler struct {
	DB      *driver.DB
	service *commonservice.NotFoundRecordService
	curd    *curd.CRUD
}

// GetRouterGroupName implements yoo.Handler.
func (n *NotFoundRecordHandler) GetRouterGroupName() string {
	return "notfound"
}

// Register implements yoo.Handler.
func (n *NotFoundRecordHandler) Register(router *yoo.RouterGroup) {
	router.WithDoc(&yoo.DocItem{
		Method: http.MethodGet,
		Path:   "",
	}).GET("", n.curd.GetListHandler(&[]model.NotFoundRecord{}))
}

var _ yoo.Handler = (*NotFoundRecordHandler)(nil)

func NewNotFoundRecordHandler(db *driver.DB, service *commonservice.NotFoundRecordService) *NotFoundRecordHandler {
	return &NotFoundRecordHandler{
		DB:      db,
		service: service,
		curd: &curd.CRUD{
			DB:    db,
			Model: &model.NotFoundRecord{},
		},
	}
}
