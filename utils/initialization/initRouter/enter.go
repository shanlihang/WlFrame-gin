package initRouter

import (
	medical "WlFrame-gin/app/medical/router"
	system "WlFrame-gin/app/system/router"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	system.SystemRouter(engine)
	medical.MedicalRouter(engine)
}
