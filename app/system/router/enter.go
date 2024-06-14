package router

import (
	"WlFrame-gin/app/system/server"
	"github.com/gin-gonic/gin"
)

func SystemRouter(e *gin.Engine) {

	// 用户模块 api
	//system := e.Group("/api/v1/system",authentication.Rbac())
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
	//systemPermission := e.Group("/api/v1/system", authentication.Rbac())
	{
		system.GET("/permission/common", server.SelectPermissionList)
		system.GET("/permission/menu", server.QueryPermissionList)
		//systemPermission.GET("/permission/tree", server.QueryMenusList)
		system.GET("/permission/tree", server.QueryMenusList)
		system.POST("/permission/add", server.AddPermission)
		system.GET("/permission/:id", server.QueryPermissionById)
		system.PUT("/permission/change", server.ChangePermission)
		system.DELETE("/permission/:id", server.RemovePermission)
	}
	systemSys := e.Group("/api/v1/system")
	// 登录注册模块 api
	{
		systemSys.POST("/sys/login", server.LoginSys)
		systemSys.POST("/sys/register", server.UserRegister)
	}
}
