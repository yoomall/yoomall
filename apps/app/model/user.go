package model

import (
	"lazyfury.github.com/yoomall-server/core"
)

type User struct {
	*core.Model
	UserName string `gorm:"not null;unique;index;column:username" json:"username"`
	Password string `gorm:"column:password" json:"-"`
	Role     int    `gorm:"column:role" json:"role"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Email    string `gorm:"not null;unique;index;column:email;validation:email" json:"email"`
	Phone    string `gorm:"not null;index;unique;column:phone;validation:phone" json:"phone"`

	Bio   string   `gorm:"column:bio" json:"bio"`
	ExtId uint     `gorm:"column:ext_id;null" json:"ext_id"`
	Ext   *UserExt `gorm:"foreignKey:ext_id;references:id;null" json:"ext"`
}

type UserExt struct {
	*core.Model
	ThridPartyId string `gorm:"column:thrid_party_id" json:"thrid_party_id"`
}

func (m *User) TableName() string {
	return "users"
}
