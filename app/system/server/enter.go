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
	println(user)
	result := dao.UpdateUser(user)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}
