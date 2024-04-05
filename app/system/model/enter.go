package model

import "gorm.io/gorm"

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
