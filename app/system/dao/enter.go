package dao

import (
	"WlFrame-gin/app/system/model"
	"WlFrame-gin/utils/global"
	"gorm.io/gorm"
)

func InsertUser() *gorm.DB {
	user := &model.SysUser{
		Name:     "小红",
		Sex:      1,
		Birthday: "2011-01-25",
		Phone:    "15945781213",
		Email:    "xiaohong@126.com",
	}
	result := global.DB.Model(model.SysUser{}).Create(user)
	return result
}
