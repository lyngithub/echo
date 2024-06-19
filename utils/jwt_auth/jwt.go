package jwt_auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	EXPIRE time.Duration
	SECRET string
)

type JWTClaims struct {
	Uuid     int64  `json:"uuid"`
	Loca     string `json:"loca"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// / GenToken
func GenToken(claims *JWTClaims) (string, error) {
	c := JWTClaims{
		claims.Uuid,
		claims.Loca,
		claims.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(EXPIRE).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(SECRET))
}

// ParseToken
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return []byte(SECRET), nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
