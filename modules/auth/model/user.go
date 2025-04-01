package model

import (
	"encoding/json"

	"github.com/lazyfury/pulse/framework"
	"github.com/lazyfury/pulse/framework/types"
	"github.com/lazyfury/pulse/helper/utils"
)

type User struct {
	*framework.Model
	UniID    string `gorm:"not null;unique;index;column:uni_id" json:"uni_id"`
	UserName string `gorm:"not null;unique;index;column:username" json:"username"`
	Password string `gorm:"column:password" json:"-"`
	Role     int    `gorm:"column:role" json:"role"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Email    string `gorm:"not null;unique;index;column:email;validation:email" json:"email"`
	Phone    string `gorm:"not null;index;unique;column:phone;validation:phone" json:"phone"`

	Bio string `gorm:"column:bio" json:"bio"`

	LastLoginAt types.LocalTime `gorm:"column:last_login_at" json:"last_login_at"`
}

var _ types.MarshalJSON = (*User)(nil)

func (m *User) TableName() string {
	return "users"
}

func (m *User) MarshalJSON() ([]byte, error) {
	type Alias User
	var a = &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	a.Phone = utils.StringUtils.HiddenPhone(m.Phone)
	a.Email = utils.StringUtils.HiddenEmail(m.Email)
	return json.Marshal(a)
}
