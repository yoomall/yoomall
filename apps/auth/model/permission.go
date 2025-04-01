package model

import (
	"github.com/lazyfury/pulse/framework"
)

type Permission struct {
	framework.Model
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
	Code string `gorm:"column:code;type:varchar(255)" json:"code"`
}

func (p *Permission) TableName() string {
	return "user_permissions"
}

type RolePermissionRef struct {
	*framework.Model
	PermissionId string      `json:"permission_id" gorm:"column:permission_id"`
	Permission   *Permission `json:"permission" gorm:"foreignKey:permission_id;references:ID;delete:SET NULL;default:null"`
	RoleId       string      `json:"role_id" gorm:"column:role_id"`
	Role         *UserRole   `json:"role" gorm:"foreignKey:role_id;references:ID;delete:SET NULL;default:null"`
}

func (p *RolePermissionRef) TableName() string {
	return "user_role_permission_refs"
}

type UserPermissionRef struct {
	*framework.Model
	UserId       string      `json:"user_id" gorm:"column:user_id"`
	User         *User       `json:"user" gorm:"foreignKey:user_id;references:ID;delete:SET NULL;default:null"`
	PermissionId string      `json:"permission_id" gorm:"column:permission_id"`
	Permission   *Permission `json:"permission" gorm:"foreignKey:permission_id;references:ID;delete:SET NULL;default:null"`
}

func (p *UserPermissionRef) TableName() string {
	return "user_permission_refs"
}
