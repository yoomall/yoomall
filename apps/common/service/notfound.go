package commonservice

import (
	"encoding/json"
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
	b, err := json.Marshal(map[string]any{
		"path":    path,
		"headers": request.Header,
		"query":   request.URL.Query(),
		"body":    request.Body,
		"origin":  request.RemoteAddr,
		"method":  request.Method,
	})
	if err != nil {
		return err
	}
	return n.DB.Create(&model.NotFoundRecord{Path: path, Request: (string)(b)}).Error
}

func (n *NotFoundRecordService) List() ([]*model.NotFoundRecord, error) {
	var res []*model.NotFoundRecord
	err := n.DB.Model(&model.NotFoundRecord{}).Find(&res).Error
	return res, err
}
