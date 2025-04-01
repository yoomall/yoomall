package model

import (
	"time"

	"github.com/lazyfury/pulse/framework"
)

type UserToken struct {
	*framework.Model
	UserId     uint      `gorm:"index" json:"user_id"`
	User       *User     `json:"user" gorm:"foreignKey:user_id;references:ID;delete:SET NULL;default:null"`
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
	IP         string    `json:"ip"`
	Agent      string    `json:"agent"`
	Device     string    `json:"device"`
	OS         string    `json:"os"`
	Browser    string    `json:"browser"`
	Hash       string    `json:"hash"` //由 ip agent device os browser 拼接hash而成
}

func (m *UserToken) TableName() string {
	return "user_tokens"
}
