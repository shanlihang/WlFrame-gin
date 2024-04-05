package router

import "github.com/gin-gonic/gin"

func SystemRouter(e *gin.Engine) {
	// 用户模块 api
	system := e.Group("/api/v1/system")
	{
		system.GET("/add", func(context *gin.Context) {
			context.String(200, "访问了 /add")
		})
	}
}
