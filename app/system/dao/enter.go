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

// 删除用户
func DeleteUser(id int64) *gorm.DB {
	result := global.DB.Delete(&model.SysUser{}, id)
	return result
}

// 新增用户组
func InsertUserGroup(group model.SysUserGroup) *gorm.DB {
	result := global.DB.Model(model.SysUserGroup{}).Create(&group)
	return result
}

// 删除用户组
func DeleteUserGroup(id int64) *gorm.DB {
	result := global.DB.Delete(&model.SysUserGroup{}, id)
	return result
}

// 修改用户组
func UpdateUserGroup(group model.SysUserGroup) *gorm.DB {
	result := global.DB.Save(&group)
	return result
}

// 查询用户组列表
func SelectUserGroupList() ([]model.SysUserGroup, *gorm.DB) {
	var groups []model.SysUserGroup
	result := global.DB.Model(model.SysUserGroup{}).Find(&groups)
	return groups, result
}

// 根据id查询用户组
func SelectUserGroupById(id int64) (model.SysUserGroup, *gorm.DB) {
	var group model.SysUserGroup
	result := global.DB.Where("ID = ?", id).First(&group)
	return group, result
}
