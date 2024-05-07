package server

import (
	"WlFrame-gin/app/system/dao"
	"WlFrame-gin/app/system/model"
	"WlFrame-gin/utils/jwt"
	"WlFrame-gin/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

// 新增用户
func AddUser(context *gin.Context) {
	user := model.SysUser{}
	if err := context.ShouldBindJSON(&user); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误信息为：%v", err))
	}

	user.Password = passwordHash(user.Password)
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

// 新增角色
func AddRole(context *gin.Context) {
	role := model.SysRole{}
	if err := context.ShouldBindJSON(&role); err != nil {
		panic(fmt.Sprintf("role数据绑定失败，错误信息为：%v", err))
	}
	result := dao.InsertRole(role)
	if result.RowsAffected != 0 {
		response.ResponseDML(context, result.RowsAffected, result.Error)
	}
}

// 删除角色
func RemoveRole(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteRole(id)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 更新角色
func ChangeRole(context *gin.Context) {
	role := model.SysRole{}
	if err := context.ShouldBindJSON(&role); err != nil {
		panic(fmt.Sprintf("role数据绑定失败，错误信息为：%v", err))
	}
	result := dao.UpdateRole(role)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 查询角色列表
func QueryRoleList(context *gin.Context) {
	role, result := dao.SelectRoleList()
	response.ResponseDQL(context, role, result.RowsAffected, result.RowsAffected, result.Error)
}

// 根据id查询角色
func QueryRoleById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	role, result := dao.SelectRoleById(id)
	response.ResponseDQL(context, role, result.RowsAffected, result.RowsAffected, result.Error)
}

// 新增权限
func AddPermission(context *gin.Context) {
	permission := model.SysPermission{}
	if err := context.ShouldBindJSON(&permission); err != nil {
		panic(fmt.Sprintf("permission数据绑定失败，错误信息为：%v", err))
	}
	result := dao.InsertPermission(permission)
	if result.RowsAffected != 0 {
		response.ResponseDML(context, result.RowsAffected, result.Error)
	}
}

// 删除权限
func RemovePermission(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeletePermission(id)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 更新权限
func ChangePermission(context *gin.Context) {
	permission := model.SysPermission{}
	if err := context.ShouldBindJSON(&permission); err != nil {
		panic(fmt.Sprintf("permission数据绑定失败，错误信息为：%v", err))
	}
	result := dao.UpdatePermission(permission)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}

// 查询权限列表
func QueryPermissionList(context *gin.Context) {
	top, result := dao.SelectTopPermission()
	for _, item := range top {
		child, _ := dao.SelectSubPermission(item.ID)
		println(child)
	}
	response.ResponseDQL(context, top, result.RowsAffected, result.RowsAffected, result.Error)
}

// 根据id查询权限
func QueryPermissionById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	role, result := dao.SelectPermissionById(id)
	response.ResponseDQL(context, role, result.RowsAffected, result.RowsAffected, result.Error)
}

// 系统登录
func LoginSys(context *gin.Context) {
	loginMsg := &model.LoginGin{}
	if err := context.ShouldBindJSON(loginMsg); err != nil {
		panic(fmt.Sprintf("LoginSys数据绑定失败，错误原因：%v", err))
	}
	result, row := dao.SelectUserAndPass(loginMsg.Username)
	if row == 0 {
		response.ResponseText(context, "账号不存在")
		return
	}

	err := ComparePassword(result.Password, loginMsg.Password)
	if err != nil {
		panic(fmt.Sprintf("密码校验异常，错误信息为：%v", err))
	}
	token, err := jwt.GenerateToken(1, "slh")
	if err != nil {
		panic(fmt.Sprintf("生成token异常，错误信息为：%v", err))
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"token": token,
	})

}

// 密码加密
func passwordHash(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(fmt.Sprintf("密码加密失败,错误原因为：%v", err))
	}
	return string(hash)
}

// 比较密码是否相同
func ComparePassword(hashPassword, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword))
	return err
}

// 注册
func UserRegister(context *gin.Context) {
	user := &model.SysUser{}
	if err := context.ShouldBindJSON(user); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误原因：%v", err))
	}
	row := dao.SelectUserByUserName(user.Username)
	if row != 0 {
		response.ResponseText(context, "用户名已存在")
		return
	}
	res := dao.InsertUser(*user)
	if res.RowsAffected != 0 {
		response.ResponseText(context, "注册成功")
	} else {
		response.ResponseText(context, "注册失败")
	}
}

// 新增验证码
func AddCaptcha(context *gin.Context) {
	c := model.SysCaptcha{}
	if err := context.ShouldBindJSON(&c); err != nil {
		panic(fmt.Sprintf("Captcha数据绑定失败，错误信息为：%v", err))
	}
	result := dao.InsertCaptcha(c)
	if result.RowsAffected != 0 {
		response.ResponseDML(context, result.RowsAffected, result.Error)
	}
}

// 删除验证码
func RemoveCaptcha(context *gin.Context) {
	k := context.Param("key")
	result := dao.DeleteCaptcha(k)
	response.ResponseDML(context, result.RowsAffected, result.Error)
}
