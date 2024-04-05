package server

import (
	"WlFrame-gin/app/system/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddUser(context *gin.Context) {

	result := dao.InsertUser()
	s := fmt.Sprintf("影响行数%d,错误信息：%v", result.RowsAffected, result.Error)

	context.String(200, s)
}
