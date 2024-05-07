package router

import (
	"WlFrame-gin/app/medical/server"
	"github.com/gin-gonic/gin"
)

func MedicalRouter(e *gin.Engine) {
	system := e.Group("/api/v1/medical")
	//结果
	{
		system.POST("/result/add", server.AddResult)
		system.GET("/result/list", server.GetResultList)
		system.DELETE("/result/:id", server.DropResult)
	}
	//社区
	{
		system.POST("/community/add", server.AddCommunity)
		system.GET("/community/list", server.GetCommunityList)
		system.DELETE("/community/:id", server.DropCommunity)
	}
	//物品
	{
		system.POST("/goods/add", server.AddGoods)
		system.GET("/goods/list", server.GetGoodsList)
		system.DELETE("/goods/:id", server.DropGoods)
	}
}