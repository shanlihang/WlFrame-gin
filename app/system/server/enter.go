package server

import (
	"WlFrame-gin/app/system/dao"
	"WlFrame-gin/app/system/model"
	"WlFrame-gin/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 新增用户
func AddUser(context *gin.Context) {
	user := model.SysUser{}
	if err := context.ShouldBindJSON(&user); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误信息为：%v", err))
	}
	result := dao.InsertUser(user)
	if result.RowsAffected != 0 {
		response.ResponseDML(context, result.RowsAffected, result.Error)
	}
}

// 查询用户列表
func QueryUserList(context *gin.Context) {
	users, result := dao.SelectUserList()
	response.ResponseDQL(context, users, result.RowsAffected, result.RowsAffected, result.Error)
}

// 根据id查询用户
func QueryUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	user, result := dao.SelectUserById(id)
	response.ResponseDQL(context, user, result.RowsAffected, result.RowsAffected, result.Error)
}

// 更新用户
func ChangeUser(context *gin.Context) {
	user := model.SysUser{}
	if err := context.ShouldBindJSON(&user); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误信息为：%v", err))
	}
	result := dao.UpdateUser(user)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 删除用户
func RemoveUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteUser(id)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 新增用户组
func AddUserGroup(context *gin.Context) {
	group := model.SysUserGroup{}
	if err := context.ShouldBindJSON(&group); err != nil {
		panic(fmt.Sprintf("userGroup数据绑定失败，错误信息为：%v", err))
	}
	result := dao.InsertUserGroup(group)
	if result.RowsAffected != 0 {
		response.ResponseDML(context, result.RowsAffected, result.Error)
	}
}

// 删除用户组
func RemoveUserGroup(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteUserGroup(id)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 更新用户组
func ChangeUserGroup(context *gin.Context) {
	group := model.SysUserGroup{}
	if err := context.ShouldBindJSON(&group); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误信息为：%v", err))
	}
	result := dao.UpdateUserGroup(group)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 查询用户组列表
func QueryUserGroupList(context *gin.Context) {
	groups, result := dao.SelectUserGroupList()
	response.ResponseDQL(context, groups, result.RowsAffected, result.RowsAffected, result.Error)
}

// 根据id查询用户组
func QueryUserGroupById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	group, result := dao.SelectUserGroupById(id)
	response.ResponseDQL(context, group, result.RowsAffected, result.RowsAffected, result.Error)
}
