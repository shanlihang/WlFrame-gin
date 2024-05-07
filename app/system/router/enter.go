package router

import (
	"WlFrame-gin/app/system/server"
	"github.com/gin-gonic/gin"
)

func SystemRouter(e *gin.Engine) {

	// 用户模块 api
	system := e.Group("/api/v1/system")
	{
		system.POST("/user/add", server.AddUser)
		system.GET("/user/list", server.QueryUserList)
		system.GET("/user/:id", server.QueryUserById)
		system.PUT("/user/change", server.ChangeUser)
		system.DELETE("/user/:id", server.RemoveUser)
	}
	// 角色模块 api
	{
		system.POST("/role/add", server.AddRole)
		system.GET("/role/list", server.QueryRoleList)
		system.GET("/role/:id", server.QueryRoleById)
		system.PUT("/role/change", server.ChangeRole)
		system.DELETE("/role/:id", server.RemoveRole)
	}
	// 权限模块 api
	{
		system.GET("/menu/list", server.QueryPermissionList)
		system.GET("/func/list", server.QueryFeaturesList)
		system.POST("/permission/add", server.AddPermission)
		system.GET("/permission/:id", server.QueryPermissionById)
		system.PUT("/permission/change", server.ChangePermission)
		system.DELETE("/permission/:id", server.RemovePermission)
	}
	// 登录注册模块 api
	{
		system.POST("/sys/login", server.LoginSys)
		system.POST("/sys/register", server.UserRegister)
		system.POST("/captcha/add", server.AddCaptcha)
		system.DELETE("/captcha/:key", server.RemoveCaptcha)
	}
}
