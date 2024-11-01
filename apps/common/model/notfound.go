package model

import "yoomall/core"

type NotFoundRecord struct {
	*core.Model
	Path    string `json:"path"`
	Request string `json:"request"`
}

func (n *NotFoundRecord) TableName() string {
	return "not_found_records"
}
