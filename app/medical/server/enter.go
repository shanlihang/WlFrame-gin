package server

import (
	"WlFrame-gin/app/medical/dao"
	"WlFrame-gin/app/medical/model"
	"WlFrame-gin/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddResult(ctx *gin.Context) {
	res := &model.Result{}
	if err := ctx.ShouldBindJSON(res); err != nil {
		panic("绑定失败")
	}
	result := dao.InsertResult(res)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}

func GetResultList(ctx *gin.Context) {
	role, result := dao.SelectResultsList()
	response.ResponseDQL(ctx, role, result.RowsAffected, result.RowsAffected, result.Error)
}

func DropResult(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteResult(id)
	response.ResponseDML(ctx, result.RowsAffected, result.Error)
}
