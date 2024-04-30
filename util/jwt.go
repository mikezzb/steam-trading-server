package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mikezzb/steam-trading-server/pkg/setting"
)

var jwtSecret []byte

type Claims struct {
	// mongodb user id
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userId, role string) (string, error) {
	expireTime := time.Now().Add(setting.App.JwtExpireMins)
	claims := Claims{
		userId,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    setting.App.AppName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
