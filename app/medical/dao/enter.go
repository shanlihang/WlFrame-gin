package dao

import (
	"WlFrame-gin/app/medical/model"
	"WlFrame-gin/utils/global"

	"gorm.io/gorm"
)

// 插入检测结果
func InsertResult(res *model.Result) *gorm.DB {
	result := global.DB.Create(res)
	return result
}

// 查询检测结果列表
func SelectResultsList() ([]model.Result, *gorm.DB) {
	var results []model.Result
	result := global.DB.Model(model.Result{}).Find(&results)
	return results, result
}

// 根据id查询检测结果
func SelectResultById(ID int64) (model.Result, *gorm.DB) {
	var res model.Result
	result := global.DB.Model(model.Result{}).Where("id = ?", ID).First(&res)
	return res, result
}

func DeleteResult(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Result{}, ID)
	return result
}

// 插入社区
func InsertCommunity(comm *model.Community) *gorm.DB {
	result := global.DB.Create(comm)
	return result
}

// 查询社区列表
func SelectCommunityList() ([]model.Community, *gorm.DB) {
	var comms []model.Community
	result := global.DB.Model(model.Community{}).Find(&comms)
	return comms, result
}

// 删除社区
func DeleteCommunity(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Community{}, ID)
	return result
}

// 插入物品
func InsertGoods(good *model.Goods) *gorm.DB {
	result := global.DB.Create(good)
	return result
}

// 查询物品列表
func SelectGoodsList() ([]model.Goods, *gorm.DB) {
	var goods []model.Goods
	result := global.DB.Model(model.Goods{}).Find(&goods)
	return goods, result
}

// 删除物品
func DeleteGoods(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Goods{}, ID)
	return result
}

// 插入推送
func InsertMsg(msg *model.PushMsg) *gorm.DB {
	result := global.DB.Create(msg)
	return result
}

// 查询推送列表
func SelectMsgList() ([]model.PushMsg, *gorm.DB) {
	var msg []model.PushMsg
	result := global.DB.Model(model.PushMsg{}).Find(&msg)
	return msg, result
}

// 根据id删除推送
func DeleteMsg(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.PushMsg{}, ID)
	return result
}

// 插入居民
func InsertPeople(people *model.People) *gorm.DB {
	result := global.DB.Create(people)
	return result
}

// 查询居民列表
func SelectPeoplesList() ([]model.People, *gorm.DB) {
	var peoples []model.People
	result := global.DB.Model(model.People{}).Find(&peoples)
	return peoples, result
}

// 删除居民
func DeletePeople(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.People{}, ID)
	return result
}

// 插入反馈
func InsertFeedback(feedback *model.Feedback) *gorm.DB {
	result := global.DB.Create(feedback)
	return result
}

// 查询反馈列表
func SelectFeedbacksList() ([]model.Feedback, *gorm.DB) {
	var feedbacks []model.Feedback
	result := global.DB.Model(model.Feedback{}).Find(&feedbacks)
	return feedbacks, result
}

// 删除反馈
func DeleteFeedback(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Feedback{}, ID)
	return result
}
