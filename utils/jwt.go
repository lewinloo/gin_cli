package utils

import (
	"time"

	"gin_cli/config"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte(config.JwtConf.JwtSecret)

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	RoleId   string `json:"roleId"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id uint, username string, roleId string) (string, error) {
	expireTime := time.Now().Add(time.Duration(config.JwtConf.ExpireTime) * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		RoleId:   roleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.JwtConf.Issuer,
			Subject:   config.JwtConf.Subject,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
