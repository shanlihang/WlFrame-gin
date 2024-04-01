package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	//监听端口，默认为8080
	r.Run()
}
