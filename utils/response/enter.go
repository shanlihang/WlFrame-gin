package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 封装增删改操作的数据返回格式
func ResponseDML(context *gin.Context, row int64, msg error) {
	context.JSON(http.StatusOK, gin.H{
		"rowAffect": row,
		"err":       msg,
	})
}

// 封装查询操作的数据返回格式
func ResponseDQL(context *gin.Context, data interface{}, total int64, row int64, msg error) {
	context.JSON(http.StatusOK, gin.H{
		"data":      data,
		"total":     total,
		"rowAffect": row,
		"err":       msg,
	})
}
