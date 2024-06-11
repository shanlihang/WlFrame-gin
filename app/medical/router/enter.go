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
		system.GET("/result/:id", server.GetResultById)
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
	//推送
	{
		system.POST("/push/add", server.AddPush)
		system.GET("/push/list", server.GetPushList)
		system.GET("/push/:id", server.GetPushById)
		system.DELETE("/push/:id", server.DropPush)
	}
	//居民
	{
		system.POST("/people/add", server.AddPeople)
		system.GET("/people/list", server.GetPeopleList)
		system.DELETE("/people/:id", server.DropPeople)
	}
	//反馈
	{
		system.POST("/feedback/add", server.AddFeedback)
		system.GET("/feedback/list", server.GetFeedbackList)
		system.DELETE("/feedback/:id", server.DropFeedback)
	}
}
