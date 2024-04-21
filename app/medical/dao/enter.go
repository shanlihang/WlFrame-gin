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
