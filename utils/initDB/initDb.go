package initDB

import (
	"WlFrame-gin/conf"
	"WlFrame-gin/utils/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	global.DBConfig = conf.GetDatabaseConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.DBConfig.Username, global.DBConfig.Password, global.DBConfig.Host, global.DBConfig.Port, global.DBConfig.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败, 错误信息为：" + err.Error())
	}
	global.DB = db
}
