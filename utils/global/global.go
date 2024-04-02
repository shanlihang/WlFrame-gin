package global

import (
	"WlFrame-gin/conf"
	"gorm.io/gorm"
)

// 数据库相关
var (
	DB       *gorm.DB
	DBConfig conf.Database
)
