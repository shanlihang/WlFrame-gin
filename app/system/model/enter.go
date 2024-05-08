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
	Birthday string `gorm:"column:birthday" json:"birthday"`
	Email    string `gorm:"column:email" json:"email"`
}

func (user SysUser) TableName() string {
	return "sys_user"
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
	Name      *string          `json:"name" gorm:"column:name"`
	Component *string          `json:"component" gorm:"column:component"`
	Icon      *string          `json:"icon" gorm:"column:icon"`
	ParentID  *int64           `json:"parentId" gorm:"column:parentId"`
	Router    *string          `json:"router" gorm:"column:router"`
	Sort      *int64           `json:"sort" gorm:"column:sort"`
	Type      *int64           `json:"type" gorm:"column:type"`
	URI       *string          `json:"uri" gorm:"column:uri"`
	Children  *[]SysPermission `json:"children" gorm:"-"`
}

func (permission SysPermission) TableName() string {
	return "sys_permission"
}

// JWT
type JWT struct {
	Username string `json:"username"`
}

// 登录
type LoginGin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 验证码
type SysCaptcha struct {
	ID        uint   `json:"id" gorm:"column:ID"`
	VerifyKey string `json:"key" gorm:"column:verifyKey"`
	VerifyImg string `json:"img" gorm:"column:verifyImg"`
}

func (captcha SysCaptcha) TableName() string {
	return "sys_captcha"
}
