/**
 * @Author: entere
 * @Description:
 * @File:  token.go
 * @Version: 1.0.0
 * @Date: 2019/12/16 21:50
 */

package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// tokenExpiredDate app token过期日期 30天
	jwtExpiredDate = 3600 * 24 * 30 * time.Second
	jwtSecretKey   = "XsX20LwAEeUQpCcXcmbzt0yI0x0sObUQ"
)

type Claims struct {
	UserID   string `json:"userID"`
	DeviceID string `json:"deviceID"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(userID string, deviceID string) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:   userID,
		DeviceID: deviceID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: now.Add(jwtExpiredDate).Unix(),
			Id:        userID,
			IssuedAt:  now.Unix(),
			Issuer:    "com.island.code",
			NotBefore: now.Unix(),
			Subject:   userID,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecretKey))

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
