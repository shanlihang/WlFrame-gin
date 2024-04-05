package dao

import (
	"WlFrame-gin/app/system/model"
	"WlFrame-gin/utils/global"
	"gorm.io/gorm"
)

// 新增用户
func InsertUser(user model.SysUser) *gorm.DB {
	result := global.DB.Model(model.SysUser{}).Create(&user)
	return result
}

// 查询用户列表
func SelectUserList() ([]model.SysUser, *gorm.DB) {
	var users []model.SysUser
	result := global.DB.Model(model.SysUser{}).Find(&users)
	return users, result
}

// 根据id查询用户
func SelectUserById(id int64) (model.SysUser, *gorm.DB) {
	var user model.SysUser
	result := global.DB.Where("ID = ?", id).First(&user)
	return user, result
}

// 修改用户
func UpdateUser(user model.SysUser) *gorm.DB {
	result := global.DB.Save(&user)
	return result
}
