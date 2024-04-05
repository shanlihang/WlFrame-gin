package model

import "gorm.io/gorm"

// 系统用户
type SysUser struct {
	gorm.Model
	Name     string `gorm:"column:name" json:"name"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Sex      int64  `gorm:"column:sex" json:"sex"`
	Slat     string `gorm:"column:slat" json:"slat"`
	Birthday string `gorm:"column:birthday" json:"birthday"`
	Email    string `gorm:"column:email" json:"email"`
}

func (user SysUser) TableName() string {
	return "sys_user"
}

// 系统用户组
type SysUserGroup struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name"`
	Desc     string `json:"desc" gorm:"column:desc"`
	ParentID int64  `json:"parentId" gorm:"column:parentId"`
}

func (group SysUserGroup) TableName() string {
	return "sys_user_group"
}

// 系统角色
type SysRole struct {
	gorm.Model
	Desc *string `json:"desc" gorm:"column:desc"`
	Name *string `json:"name" gorm:"column:name"`
}

func (role SysRole) TableName() string {
	return "sys_role"
}

// 系统权限
type SysPermission struct {
	gorm.Model
	Name      *string `json:"name" gorm:"column:name"`
	Component *string `json:"component" gorm:"column:component"`
	Icon      *string `json:"icon" gorm:"column:icon"`
	ParentID  *int64  `json:"parentId" gorm:"column:parentId"`
	Router    *string `json:"router" gorm:"column:router"`
	Sort      *int64  `json:"sort" gorm:"column:sort"`
	Type      *int64  `json:"type" gorm:"column:type"`
	URI       *string `json:"uri" gorm:"column:uri"`
}

func (permission SysPermission) TableName() string {
	return "sys_permission"
}
