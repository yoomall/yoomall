package handler

import (
	"net/http"

	"yoomall/apps/common/model"
	commonservice "yoomall/apps/common/service"

	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/driver"
	"github.com/lazyfury/pulse/helper/curd"
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
func (n *NotFoundRecordHandler) Register(router *framework.RouterGroup) {
	router.Doc(&framework.DocItem{
		Method: http.MethodGet,
		Path:   "",
	}).GET("", n.curd.GetListHandler(&[]model.NotFoundRecord{}))
}

var _ framework.Handler = (*NotFoundRecordHandler)(nil)

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
