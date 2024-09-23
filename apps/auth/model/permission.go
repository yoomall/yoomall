package model

import "lazyfury.github.com/yoomall-server/core"

type Permission struct {
	core.Model
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
	Slug string `gorm:"column:slug;type:varchar(255)" json:"slug"`
}

func (p *Permission) TableName() string {
	return "user_permissions"
}

type RolePermissionRef struct {
	*core.Model
	PermissionId string      `json:"permissionId" gorm:"column:permission_id"`
	Permission   *Permission `json:"permission" gorm:"foreignKey:permission_id;references:ID;delete:SET NULL;default:null"`
	RoleId       string      `json:"roleId" gorm:"column:role_id"`
	Role         *UserRole   `json:"role" gorm:"foreignKey:role_id;references:ID;delete:SET NULL;default:null"`
}

func (p *RolePermissionRef) TableName() string {
	return "user_role_permission_refs"
}

type UserPermissionRef struct {
	*core.Model
	UserId       string      `json:"userId" gorm:"column:user_id"`
	User         *User       `json:"user" gorm:"foreignKey:user_id;references:ID;delete:SET NULL;default:null"`
	PermissionId string      `json:"permissionId" gorm:"column:permission_id"`
	Permission   *Permission `json:"permission" gorm:"foreignKey:permission_id;references:ID;delete:SET NULL;default:null"`
}

func (p *UserPermissionRef) TableName() string {
	return "user_permission_refs"
}
