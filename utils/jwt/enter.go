package jwt

import (
	"WlFrame-gin/utils/global"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var jwtSecret = []byte("AllYourBase")

type Claims struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
	Role []string `json:"role"`
}

func GenerateToken(id int64, name string, role []string) (string, error) {
	claims := &Claims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: global.EffectTime,
			ExpiresAt: int64(time.Hour * 24 * 7),
			Issuer:    global.Issuer,
		},
		Role: role,
	}
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(global.Secret)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(global.Secret)
}

// 解析token
func AnalysisToken(tokenString string) (*Claims, bool) {
	// 去除tokenString中的"Bearer "前缀
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return global.Secret, nil
	})
	if err != nil {
		return &Claims{}, false
	}
	// 对token对象中的Claim进行类型断言
	if claims_, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims_, true
	}

	return &Claims{}, false
}
