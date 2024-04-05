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
	}
}
