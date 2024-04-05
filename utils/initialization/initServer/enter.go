package initServer

import (
	"WlFrame-gin/conf"
	"WlFrame-gin/utils/global"
)

func init() {
	s := conf.GetServerConfig()
	global.ServerConfig = s
}
