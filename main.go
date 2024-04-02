package main

import (
	"WlFrame-gin/utils/initDB"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由
	r := gin.Default()

	s := initDB.InitDataBaseConnection()
	//绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, s)
	})
	//监听端口，默认为8080
	r.Run()
}
