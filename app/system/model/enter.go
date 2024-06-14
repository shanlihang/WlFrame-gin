package model

import "gorm.io/gorm"

// 系统用户
type SysUser struct {
	gorm.Model
	Name     string    `gorm:"column:name" json:"name"`
	Username string    `gorm:"column:username" json:"username"`
	Password string    `gorm:"column:password" json:"password"`
	Phone    string    `gorm:"column:phone" json:"phone"`
	Sex      int64     `gorm:"column:sex" json:"sex"`
	Birthday string    `gorm:"column:birthday" json:"birthday"`
	Email    string    `gorm:"column:email" json:"email"`
	Roles    []SysRole `gorm:"many2many:relate_user_role"`
}

// 解析前端数据
type SysUserFrontEnd struct {
	gorm.Model
	Name     string  `gorm:"column:name" json:"name"`
	Username string  `gorm:"column:username" json:"username"`
	Password string  `gorm:"column:password" json:"password"`
	Phone    string  `gorm:"column:phone" json:"phone"`
	Sex      int64   `gorm:"column:sex" json:"sex"`
	Birthday string  `gorm:"column:birthday" json:"birthday"`
	Email    string  `gorm:"column:email" json:"email"`
	Roles    []int64 `json:"roles"`
}

// 用户更新接收json数据
type SysUserChange struct {
	gorm.Model
	Name     string    `gorm:"column:name" json:"name"`
	Username string    `gorm:"column:username" json:"username"`
	Password string    `gorm:"column:password" json:"password"`
	Phone    string    `gorm:"column:phone" json:"phone"`
	Sex      int64     `gorm:"column:sex" json:"sex"`
	Birthday string    `gorm:"column:birthday" json:"birthday"`
	Email    string    `gorm:"column:email" json:"email"`
	SysRoles []SysRole `gorm:"many2many:relate_user_role" json:"Roles"`
	Roles    []int64   `json:"roles"`
}

func (user SysUser) TableName() string {
	return "sys_user"
}

// 用户-角色 连接表
type RelateUserRole struct {
	SysUserID uint `gorm:"primaryKey"`
	SysRoleID uint `gorm:"primaryKey"`
}

func (userRole RelateUserRole) TableName() string {
	return "relate_User_role"
}

// 系统角色
type SysRole struct {
	gorm.Model
	Desc string `json:"desc" gorm:"column:desc"`
	Name string `json:"name" gorm:"column:name"`
	//UserID      int64           `json:"-" gorm:"many2many:user_credit_cards;"`
	Permissions []SysPermission `json:"permissions" gorm:"many2many:relate_role_permission"`
}

// 接收前端给角色设置权限
type SysRoleFrontEnd struct {
	gorm.Model
	Desc        string         `json:"desc" gorm:"column:desc"`
	Name        string         `json:"name" gorm:"column:name"`
	UserID      int64          `json:"-" gorm:"many2many:user_credit_cards;"`
	Permissions []Jurisdiction `json:"permissions" gorm:"-"`
}

// 角色的权限
type Jurisdiction struct {
	Label string `json:"label"`
	Value uint   `json:"value"`
}

func (role SysRole) TableName() string {
	return "sys_role"
}

// 角色-系统权限 连接表
type RelateRolePermission struct {
	SysRoleID       uint `gorm:"primaryKey"`
	SysPermissionID uint `gorm:"primaryKey"`
}

func (userRole RelateRolePermission) TableName() string {
	return "relate_role_permission"
}

// 系统权限
type SysPermission struct {
	gorm.Model
	Name      string           `json:"name" gorm:"column:name"`
	Component string           `json:"component" gorm:"column:component"`
	Icon      string           `json:"icon" gorm:"column:icon"`
	ParentID  int64            `json:"parentId" gorm:"column:parentId"`
	Router    string           `json:"router" gorm:"column:router"`
	Sort      int64            `json:"sort" gorm:"column:sort"`
	Type      int64            `json:"type" gorm:"column:type"`
	URI       string           `json:"uri" gorm:"column:uri"`
	Children  []*SysPermission `json:"children" gorm:"-"`
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
