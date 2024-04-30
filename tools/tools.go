package tools

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtKey = "sample"
)

type Claims struct {
	jwt.RegisteredClaims

	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func GenerateToken(uid int64, name string) (string, error) {
	// Generate jwt, 临时有效期 360 分钟
	nowTime := time.Now()
	expiresTime := nowTime.Add(360 * time.Minute)
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(nowTime),     // 签发时间
			NotBefore: jwt.NewNumericDate(nowTime),     // 生效时间
		},
		Id:   uid,
		Name: name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("准备返回处理")
	return token.SignedString([]byte(jwtKey))
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to parse invailed token")
}
