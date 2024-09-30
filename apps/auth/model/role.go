package model

import "lazyfury.github.com/yoomall-server/core"

type UserRole struct {
	*core.Model
	RoleName string `json:"roleName" gorm:"column:role_name"`
	RoleCode string `json:"roleCode" gorm:"column:role_code;unique;index:role_code;default:''"`
	RoleDesc string `json:"roleDesc" gorm:"column:role_desc;default:''"`
}

func (u *UserRole) TableName() string {
	return "user_roles"
}

type UserRoleRef struct {
	*core.Model
	UserId string    `json:"userId" gorm:"column:user_id"`
	User   *User     `json:"user" gorm:"foreignKey:user_id;references:ID;delete:SET NULL;default:null"`
	RoleId string    `json:"roleId" gorm:"column:role_id"`
	Role   *UserRole `json:"role" gorm:"foreignKey:role_id;references:ID;delete:SET NULL;default:null"`
}

func (u *UserRoleRef) TableName() string {
	return "user_role_refs"
}
