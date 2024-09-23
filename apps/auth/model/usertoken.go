package model

import (
	"time"

	"lazyfury.github.com/yoomall-server/core"
)

type UserToken struct {
	*core.Model
	UserId     uint      `gorm:"index" json:"user_id"`
	User       *User     `json:"user" gorm:"foreignKey:user_id;references:ID;delete:SET NULL;default:null"`
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

func (m *UserToken) TableName() string {
	return "user_tokens"
}
