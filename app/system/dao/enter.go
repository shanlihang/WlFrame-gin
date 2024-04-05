package dao

import (
	"WlFrame-gin/app/system/model"
	"WlFrame-gin/utils/global"
	"gorm.io/gorm"
)

func InsertUser(user *model.SysUser) *gorm.DB {
	result := global.DB.Model(model.SysUser{}).Create(user)
	return result
}
