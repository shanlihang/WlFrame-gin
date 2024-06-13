package jwt

import (
	"WlFrame-gin/utils/global"
	"github.com/dgrijalva/jwt-go"
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
			ExpiresAt: global.EffectTime,
			Issuer:    global.Issuer,
		},
		Role: role,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(global.Secret)
	return token, err
}

// 解析token
func AnalysisToken(tokenString string) (*Claims, bool) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
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
