package initRouter

import (
	"WlFrame-gin/app/system/router"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	router.SystemRouter(engine)
}
