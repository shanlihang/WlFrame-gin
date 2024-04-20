package router

import (
	"WlFrame-gin/app/medical/server"
	"github.com/gin-gonic/gin"
)

func MedicalRouter(e *gin.Engine) {
	system := e.Group("/api/v1/medical")
	{
		system.POST("/result/add", server.AddResult)
		system.GET("/result/list", server.GetResultList)
		system.DELETE("/result/:id", server.DropResult)
	}
}
