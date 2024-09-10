package model

import (
	"encoding/json"

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
	ExtId uint     `gorm:"column:ext_id;default:null" json:"ext_id"`
	Ext   *UserExt `gorm:"foreignKey:ext_id;references:id;delete:SET NULL;default:null" json:"ext"`
}

type UserExt struct {
	*core.Model
	ThridPartyId string `gorm:"column:thrid_party_id" json:"thrid_party_id"`
}

var _ core.MarshalJSON = (*User)(nil)

func (m *User) TableName() string {
	return "users"
}

func (m *User) MarshalJSON() ([]byte, error) {
	type Alias User
	var a = &struct {
		*Alias
		NewPhone string `json:"new_phone"`
	}{
		Alias:    (*Alias)(m),
		NewPhone: m.hiddenPhone(),
	}
	a.Phone = m.hiddenPhone()
	return json.Marshal(a)
}

func (m *User) hiddenPhone() string {
	phone := m.Phone
	if phone == "" {
		return ""
	}
	if len(phone) <= 4 {
		return phone
	}
	start := 3
	end := len(phone) - 4
	return phone[0:start] + "****" + phone[end:]
}
