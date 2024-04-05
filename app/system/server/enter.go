package server

import (
	"WlFrame-gin/app/system/dao"
	"WlFrame-gin/app/system/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddUser(context *gin.Context) {
	user := &model.SysUser{}
	if err := context.ShouldBindJSON(user); err != nil {
		panic(fmt.Sprintf("user数据绑定失败，错误信息为：%v", err))
	}
	result := dao.InsertUser(user)
	if result.RowsAffected != 0 {
		context.JSON(200, gin.H{
			//"code":      0,
			//"msg":       "新增成功",
			//"rowAffect": result.RowsAffected,
			"result": result,
		})
	}

}
