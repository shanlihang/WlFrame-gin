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
