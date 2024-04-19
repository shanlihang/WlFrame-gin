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

// 根据账号查询密码
func SelectUserAndPass(username string) (model.LoginGin, int64) {
	login := model.LoginGin{}
	result := global.DB.Model(model.SysUser{}).Select("username", "password").Where("username =? ", username).First(&login)
	return login, result.RowsAffected
}

// 根据账号查询密码
func SelectUserByUserName(username string) int64 {
	user := model.SysUser{}
	result := global.DB.Model(model.SysUser{}).Where("username =? ", username).First(&user)
	return result.RowsAffected
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

// 查询菜单列表
func SelectPermissionList() ([]model.SysPermission, *gorm.DB) {
	var permissions []model.SysPermission
	var subPermissions []model.SysPermission
	//查询顶级菜单
	result := global.DB.Model(model.SysPermission{}).Where("type = ?", 0).Find(&permissions)

	//查询子菜单并拼接
	for i := 0; i < len(permissions); i++ {
		_ = global.DB.Model(model.SysPermission{}).Where("parentId = ?", permissions[i].ID).Find(&subPermissions)
		permissions[i].Children = subPermissions
	}

	return permissions, result
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

// 新增主题
func InsertTheme(theme model.SysTheme) *gorm.DB {
	result := global.DB.Model(model.SysTheme{}).Create(&theme)
	return result
}

// 查询主题列表
func SelectThemeList() ([]model.SysTheme, *gorm.DB) {
	var themes []model.SysTheme
	result := global.DB.Model(model.SysTheme{}).Find(&themes)
	return themes, result
}

// 根据id查询主题
func SelectThemeById(id int64) (model.SysTheme, *gorm.DB) {
	var theme model.SysTheme
	result := global.DB.Where("ID = ?", id).First(&theme)
	return theme, result
}

// 修改主题
func UpdateTheme(theme model.SysTheme) *gorm.DB {
	result := global.DB.Save(&theme)
	return result
}

// 删除主题
func DeleteTheme(id int64) *gorm.DB {
	result := global.DB.Delete(&model.SysTheme{}, id)
	return result
}
