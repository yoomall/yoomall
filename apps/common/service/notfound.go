package commonservice

import (
	"net/http"

	"lazyfury.github.com/yoomall-server/apps/common/model"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type NotFoundRecordService struct {
	DB *driver.DB
}

func NewNotFoundRecordService(db *driver.DB) *NotFoundRecordService {
	return &NotFoundRecordService{
		DB: db,
	}
}

func (n *NotFoundRecordService) Add(path string, request *http.Request) error {
	return n.DB.Create(&model.NotFoundRecord{Path: path}).Error
}
