package server

import (
	"WlFrame-gin/app/medical/dao"
	"WlFrame-gin/app/medical/model"
	"WlFrame-gin/utils/response"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 结果
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
	examNo := ctx.Query("examNo")
	deviceID := ctx.Query("deviceID")
	name := ctx.Query("name")
	idnumber := ctx.Query("idnumber")
	result, err := dao.SelectResultsList(examNo, deviceID, name, idnumber)
	ctx.JSON(200, gin.H{
		"data":   result,
		"errMsg": err,
	})

}
func GetResultById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	res, result := dao.SelectResultById(id)
	response.ResponseDQL(ctx, res, result.RowsAffected, result.RowsAffected, result.Error)
}
func DropResult(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteResult(id)
	response.ResponseDML(ctx, result.RowsAffected, result.Error)
}

// 社区
func AddCommunity(ctx *gin.Context) {
	res := &model.Community{}
	if err := ctx.ShouldBindJSON(res); err != nil {
		panic("绑定失败")
	}
	result := dao.InsertCommunity(res)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}
func GetCommunityList(ctx *gin.Context) {
	name := ctx.Query("name")
	district := ctx.Query("district")
	address := ctx.Query("detail_address")
	list, err := dao.SelectCommunityList(name, district, address)
	ctx.JSON(200, gin.H{
		"data":   list,
		"errMsg": err,
	})
}
func DropCommunity(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteCommunity(id)
	response.ResponseDML(ctx, result.RowsAffected, result.Error)
}

// 物品
func AddGoods(ctx *gin.Context) {
	goods := &model.Goods{}
	if err := ctx.ShouldBindJSON(goods); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	result := dao.InsertGoods(goods)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}
func GetGoodsList(ctx *gin.Context) {
	name := ctx.Query("name")
	remark := ctx.Query("remark")
	goods, err := dao.SelectGoodsList(name, remark)
	ctx.JSON(200, gin.H{
		"data":   goods,
		"errMsg": err,
	})
}
func PutGood(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err1 != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err1))
	}
	num, err2 := strconv.ParseInt(ctx.Query("num"), 10, 64)
	if err2 != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err2))
	}
	good, _ := dao.SelectGoodsById(id)
	good.Num += num
	res := dao.UpdateGood(&good)
	response.ResponseDML(ctx, res.RowsAffected, res.Error)
}
func OutGood(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err1 != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err1))
	}
	num, err2 := strconv.ParseInt(ctx.Query("num"), 10, 64)
	if err2 != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err2))
	}
	good, _ := dao.SelectGoodsById(id)
	if good.Num < num {
		response.ResponseDML(ctx, 0, errors.New("出库失败，库存数量少于提取数量"))
	} else {
		good.Num -= num
		res := dao.UpdateGood(&good)
		response.ResponseDML(ctx, res.RowsAffected, res.Error)
	}
}
func GetGoodById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	user, result := dao.SelectGoodsById(id)
	response.ResponseDQL(ctx, user, result.RowsAffected, result.RowsAffected, result.Error)
}
func DropGoods(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteGoods(id)
	response.ResponseDML(ctx, result.RowsAffected, result.Error)
}
func ChangeGoods(ctx *gin.Context) {
	good := &model.Goods{}
	if err := ctx.ShouldBindJSON(good); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	fmt.Println(good.ID)
	result := dao.UpdateGood(good)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}

// 推送
func AddPush(ctx *gin.Context) {
	msg := &model.PushMsg{}
	if err := ctx.ShouldBindJSON(msg); err != nil {
		panic("绑定失败")
	}
	result := dao.InsertMsg(msg)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}
func GetPushList(ctx *gin.Context) {
	title := ctx.Query("title")
	content := ctx.Query("content")
	list, err := dao.SelectMsgList(title, content)
	ctx.JSON(200, gin.H{
		"data":   list,
		"errMsg": err,
	})
}
func GetPushById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	res, result := dao.SelectMsgById(id)
	response.ResponseDQL(ctx, res, result.RowsAffected, result.RowsAffected, result.Error)
}
func DropPush(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteMsg(id)
	response.ResponseDML(ctx, result.RowsAffected, result.Error)
}
func UpdatePush(ctx *gin.Context) {
	msg := &model.PushMsg{}
	if err := ctx.ShouldBindJSON(msg); err != nil {
		panic("绑定失败")
	}
	result := dao.UpdateMsg(msg)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}

// 居民
func AddPeople(ctx *gin.Context) {
	people := &model.People{}
	if err := ctx.ShouldBindJSON(people); err != nil {
		panic("绑定失败")
	}
	result := dao.InsertPeople(people)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}
func GetPeopleList(ctx *gin.Context) {
	name := ctx.Query("name")
	email := ctx.Query("email")
	phone := ctx.Query("phone")
	idnumber := ctx.Query("idnumber")

	list, err := dao.SelectPeoplesList(name, email, phone, idnumber)
	ctx.JSON(200, gin.H{
		"data":   list,
		"errMsg": err,
	})
}
func DropPeople(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeletePeople(id)
	response.ResponseDML(ctx, result.RowsAffected, result.Error)
}
func UpdatePeople(ctx *gin.Context) {
	p := model.People{}
	if err := ctx.ShouldBindJSON(&p); err != nil {
		panic("绑定失败")
	}
	result := dao.UpdatePeople(p)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}

// 反馈
func AddFeedback(ctx *gin.Context) {
	feedback := &model.Feedback{}
	if err := ctx.ShouldBindJSON(feedback); err != nil {
		panic("绑定失败")
	}
	result := dao.InsertFeedback(feedback)
	if result.RowsAffected != 0 {
		response.ResponseDML(ctx, result.RowsAffected, result.Error)
	}
}
func GetFeedbackList(ctx *gin.Context) {
	content := ctx.Query("content")
	status := ctx.Query("status")
	list, err := dao.SelectFeedbacksList(content, status)
	ctx.JSON(200, gin.H{
		"data":   list,
		"errMsg": err,
	})
}
func DropFeedback(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	result := dao.DeleteFeedback(id)
	response.ResponseDML(ctx, result.RowsAffected, result.Error)
}
func ChangeFeedback(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("id属性转换为int64类型失败，错误原因：%v", err))
	}
	res := dao.UpdateFeedback(id)
	response.ResponseDML(ctx, res.RowsAffected, res.Error)
}
