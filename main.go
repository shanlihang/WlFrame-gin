package main

import (
	_ "WlFrame-gin/utils/initDB"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由
	r := gin.Default()

	//绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "水水水水水水水水水水水水水水")
	})
	////监听端口，默认为8080
	r.Run()
}
