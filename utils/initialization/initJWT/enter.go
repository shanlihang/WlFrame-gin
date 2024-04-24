package initJWT

import "WlFrame-gin/utils/global"

func init() {
	global.Secret = []byte("slh67490009")
	global.EffectTime = 86400
	global.Issuer = "slh"
}
