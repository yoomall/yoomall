package model

import (
	"time"

	"lazyfury.github.com/yoomall-server/core"
)

type UserToken struct {
	*core.Model
	UserId     uint      `gorm:"index" json:"user_id"`
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

func (m *UserToken) TableName() string {
	return "user_tokens"
}
