package model

import "github.com/lazyfury/pulse/framework"

type NotFoundRecord struct {
	*framework.Model
	Path    string `json:"path"`
	Request string `json:"request"`
}

func (n *NotFoundRecord) TableName() string {
	return "not_found_records"
}
