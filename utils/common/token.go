package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("general secret key")

type Claims struct {
	jwt.StandardClaims
	UserID   int    `json:"id"`
	Username string `json:"username"`
}

// GenerateToken 颁发token
func GenerateToken(id int, username string) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)
	claims := Claims{
		UserID:   id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

// ParseToken 验证用户token
func ParseToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
