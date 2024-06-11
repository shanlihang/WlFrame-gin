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
func SelectUserList(name string, phone string, email string, roles []string) ([]model.SysUser, error) {
	var users []model.SysUser
	tx := global.DB.Model(model.SysUser{}).Omit("Password")
	if name != "" {
		tx.Where("name Like ?", "%"+name+"%")
	}
	if phone != "" {
		tx.Where("phone Like ?", "%"+phone+"%")
	}
	if email != "" {
		tx.Where("email Like ?", "%"+email+"%")
	}
	if len(roles) != 0 {
		tx.Preload("Roles", "ID in ?", roles)
	}
	err := tx.Preload("Roles").Find(&users).Error
	return users, err
}

// 根据id查询用户
func SelectUserById(id int64) (model.SysUser, *gorm.DB) {
	var user model.SysUser
	result := global.DB.Where("ID = ?", id).Omit("Password").First(&user)
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

// 根据账号查询密码
func SelectUserAndPass(username string) (model.LoginGin, int64) {
	login := model.LoginGin{}
	result := global.DB.Model(model.SysUser{}).Select("username", "password").Where("username =? ", username).First(&login)
	return login, result.RowsAffected
}

// 根据账号查询用户
func SelectUserByUserName(username string) int64 {
	user := model.SysUser{}
	result := global.DB.Model(model.SysUser{}).Where("username =? ", username).First(&user)
	return result.RowsAffected
}

// 新增角色
func InsertRole(role model.SysRole) *gorm.DB {
	result := global.DB.Model(model.SysRole{}).Create(&role)
	return result
}

// 删除角色
func DeleteRole(id int64) *gorm.DB {
	result := global.DB.Delete(&model.SysRole{}, id)
	return result
}

// 修改角色
func UpdateRole(role model.SysRole) *gorm.DB {
	result := global.DB.Save(&role)
	return result
}

// 查询角色列表
func SelectRoleList() ([]model.SysRole, *gorm.DB) {
	var roles []model.SysRole
	result := global.DB.Model(model.SysRole{}).Find(&roles)
	return roles, result
}

// 根据id查询角色
func SelectRoleById(id int64) (model.SysRole, *gorm.DB) {
	var role model.SysRole
	result := global.DB.Where("ID = ?", id).First(&role)
	return role, result
}

// 新增权限
func InsertPermission(permission model.SysPermission) *gorm.DB {
	result := global.DB.Model(model.SysPermission{}).Create(&permission)
	return result
}

// 删除权限
func DeletePermission(id int64) *gorm.DB {
	result := global.DB.Delete(&model.SysPermission{}, id)
	return result
}

// 修改权限
func UpdatePermission(permission model.SysPermission) *gorm.DB {
	result := global.DB.Save(&permission)
	return result
}

// 查询顶级菜单列表
func SelectTopPermission() ([]model.SysPermission, *gorm.DB) {
	var permissions []model.SysPermission
	result := global.DB.Model(model.SysPermission{}).Where("type = ?", 0).Find(&permissions)
	return permissions, result
}

// 查询子级菜单列表
func SelectSubPermission(id uint) ([]model.SysPermission, *gorm.DB) {
	var permissions []model.SysPermission
	result := global.DB.Model(model.SysPermission{}).Where("parentId = ? ", id).Find(&permissions)
	return permissions, result
}

// 查询目录列表
func SelectDirectory() ([]model.SysPermission, *gorm.DB) {
	var directory []model.SysPermission
	result := global.DB.Model(model.SysPermission{}).Where("type = ?", 1).Find(&directory)
	return directory, result
}

// 查询目录下的功能
func SelectFeatures(id uint) ([]model.SysPermission, *gorm.DB) {
	var features []model.SysPermission
	result := global.DB.Model(model.SysPermission{}).Where("parentId = ?", id).Find(&features)
	return features, result
}

// 根据id查询权限
func SelectPermissionById(id int64) (model.SysPermission, *gorm.DB) {
	var permission model.SysPermission
	result := global.DB.Where("ID = ?", id).First(&permission)
	return permission, result
}

// 新增验证码
func InsertCaptcha(captcha model.SysCaptcha) *gorm.DB {
	result := global.DB.Model(model.SysCaptcha{}).Create(&captcha)
	return result
}

// 删除验证码
func DeleteCaptcha(key string) *gorm.DB {
	result := global.DB.Where("verifyKey = ?", key).Delete(&model.SysCaptcha{})
	return result
}
