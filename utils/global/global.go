package global

import (
	"WlFrame-gin/conf"
	"gorm.io/gorm"
)

// 配置文件相关
var (
	DBConfig     conf.Database
	ServerConfig conf.Server
)

// 数据库相关
var (
	DB *gorm.DB
)

// JWT相关
var (
	Secret     []byte //签名信息
	EffectTime int64  //有效时长（s）
	Issuer     string //签发人
)
