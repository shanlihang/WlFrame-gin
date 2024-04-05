package main

import (
	"WlFrame-gin/utils/global"
	_ "WlFrame-gin/utils/initialization/initDb"
	"WlFrame-gin/utils/initialization/initRouter"
	_ "WlFrame-gin/utils/initialization/initServer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由
	r := gin.Default()

	//绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "水水水火火水水水水水")
	})

	initRouter.InitRouter(r)

	////监听端口，默认为8080
	if err := r.Run(":" + global.ServerConfig.Port); err == nil {
		panic(fmt.Sprintf("服务启动失败,失败原因：%v", err))
	}
}
