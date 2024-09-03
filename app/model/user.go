package model

type User struct {
	UserName string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Role     int    `gorm:"column:role" json:"role"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Email    string `gorm:"column:email" json:"email"`
	Phone    string `gorm:"column:phone" json:"phone"`

	Bio string `gorm:"column:bio" json:"bio"`
}

func (m *User) TableName() string {
	return "users"
}
