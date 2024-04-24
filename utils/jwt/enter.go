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
}

func GenerateToken(id int64, name string) (string, error) {
	claims := &Claims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: global.EffectTime,
			Issuer:    global.Issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(global.Secret)
	return token, err
}
