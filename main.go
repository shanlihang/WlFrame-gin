package main

import (
	"WlFrame-gin/middleware/cors"
	"WlFrame-gin/utils/global"
	_ "WlFrame-gin/utils/initialization/initDb"
	_ "WlFrame-gin/utils/initialization/initJWT"
	"WlFrame-gin/utils/initialization/initRouter"
	_ "WlFrame-gin/utils/initialization/initServer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由
	r := gin.Default()
	r.Use(cors.Cors())

	initRouter.InitRouter(r)

	////监听端口，默认为8080
	if err := r.Run(":" + global.ServerConfig.Port); err == nil {
		panic(fmt.Sprintf("服务启动失败,失败原因：%v", err))
	}
}
