package model

import "lazyfury.github.com/yoomall-server/core"

type UserRole struct {
	*core.Model
	RoleName string `json:"roleName" gorm:"column:role_name"`
	RoleCode string `json:"roleCode" gorm:"column:role_code"`
}

func (u *UserRole) TableName() string {
	return "user_roles"
}

type UserRoleRef struct {
	*core.Model
	UserId string `json:"userId" gorm:"column:user_id"`
	RoleId string `json:"roleId" gorm:"column:role_id"`
}

func (u *UserRoleRef) TableName() string {
	return "user_role_refs"
}
