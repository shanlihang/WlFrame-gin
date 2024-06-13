package server

import (
	"WlFrame-gin/app/system/dao"
	"WlFrame-gin/app/system/model"
	"WlFrame-gin/utils/authentication"
	"WlFrame-gin/utils/jwt"
	"WlFrame-gin/utils/response"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 新增用户
func AddUser(context *gin.Context) {
	//user := model.SysUser{}
	user := model.SysUserFrontEnd{}
	if err := context.ShouldBindJSON(&user); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误信息为：%v", err))
	}
	user.Password = passwordHash(user.Password)
	//创建用户
	result := dao.InsertUser(model.SysUser{
		Model:    user.Model,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Sex:      user.Sex,
		Birthday: user.Birthday,
		Email:    user.Email,
	}, user.Roles)

	if result.RowsAffected != 0 {
		response.ResponseDML(context, result.RowsAffected, result.Error)
	}
}

// 查询用户列表
func QueryUserList(context *gin.Context) {
	name := context.Query("name")
	phone := context.Query("phone")
	email := context.Query("email")
	roles := context.QueryArray("roles[]")
	users, err := dao.SelectUserList(name, phone, email, roles)
	context.JSON(200, gin.H{
		"data":   users,
		"errMsg": err,
	})
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
	//user := model.SysUser{}
	user := model.SysUserFrontEnd{}
	if err := context.ShouldBindJSON(&user); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误信息为：%v", err))
	}
	result := dao.UpdateUser(model.SysUser{
		Model:    user.Model,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Sex:      user.Sex,
		Birthday: user.Birthday,
		Email:    user.Email,
	}, user.Roles)
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
	role := model.SysRoleFrontEnd{}
	if err := context.BindJSON(&role); err != nil {
		panic(fmt.Sprintf("role数据绑定失败，错误信息为：%v", err))
	}
	// 根据权限id查询权限信息
	for _, permission := range role.Permissions {
		sysPermission, _ := dao.SelectPermissionById(int64(permission.Value))
		//增加 Policy(角色的权限；此次获取不了请求方法，全部设置为 GET，权限校验时不校验请求方法)
		ok := authentication.Enforcer.AddPolicy(role.Name, sysPermission.URI, "GET")
		if !ok {
			fmt.Println("Policy已经存在")
		}
	}

	result := dao.InsertRole(model.SysRole{
		Model: role.Model,
		Desc:  role.Desc,
		Name:  role.Name,
		//UserID: role.UserID,
	}, role.Permissions)
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
	//role := model.SysRole{}
	role := model.SysRoleFrontEnd{}
	if err := context.ShouldBindJSON(&role); err != nil {
		panic(fmt.Sprintf("role数据绑定失败，错误信息为：%v", err))
	}
	result := dao.UpdateRole(model.SysRole{
		Model: role.Model,
		Desc:  role.Desc,
		Name:  role.Name,
	}, role.Permissions)
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

// 查询树状菜单列表
func QueryPermissionList(context *gin.Context) {
	top, result := dao.SelectTopPermission()
	for _, item := range top {
		child, _ := dao.SelectSubPermission(item.ID)
		item.Children = append(item.Children, child...)
	}
	response.ResponseDQL(context, top, result.RowsAffected, result.RowsAffected, result.Error)
}

// 查询树状功能列表
func QueryMenusList(context *gin.Context) {
	top, result := dao.SelectTopPermission()
	for _, item := range top {
		child, _ := dao.SelectSubPermission(item.ID)
		item.Children = append(item.Children, child...)
		for _, i := range child {
			childs, _ := dao.SelectSubPermission(i.ID)
			i.Children = append(i.Children, childs...)
		}
	}
	response.ResponseDQL(context, top, result.RowsAffected, result.RowsAffected, result.Error)
}

// 查询普通列表
func SelectPermissionList(context *gin.Context) {
	list, result := dao.SelectDirectory()
	response.ResponseDQL(context, list, result.RowsAffected, result.RowsAffected, result.Error)
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
		response.ResponseText(context, "密码错误")
		context.Next()
		return
	}
	roleNames, _ := dao.SelectRelateUserRoleById(int64(result.ID))
	token, err := jwt.GenerateToken(1, "slh", roleNames)
	if err != nil {
		response.ResponseText(context, "token生成出错")
		return
	}
	result.Password = ""
	log.Println("token:", token)
	context.JSON(http.StatusOK, gin.H{
		"msg":      "登录成功",
		"token":    token,
		"userInfo": result,
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
	res := dao.InsertUser(*user, []uint{}) //注册用户时。角色为空
	if res.RowsAffected != 0 {
		response.ResponseText(context, "注册成功")
	} else {
		response.ResponseText(context, "注册失败")
	}
}
