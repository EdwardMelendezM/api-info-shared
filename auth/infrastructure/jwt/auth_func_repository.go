package jwt

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	tokenSecret = os.Getenv("JWT_SECRET")
)

func (a authJwtRepo) GenerateToken(userId string) (*string, error) {
	claims := jwt.MapClaims{
		"iss": "http://macsalud-v2.stg.erp.onscp.com/api/core/usuarios/login",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
		"nbf": time.Now().Unix(),
		"jti": "EhTxE4Iu6Is0QsBp",
		"sub": userId,
		"prv": "23bd5c8949f600adb39e701c400872db7a5976f7",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (a authJwtRepo) DecodeToken(context context.Context, tokenStr string) (userId *string, err error) {
	var token *jwt.Token

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err = jwt.ErrSignatureInvalid
			return nil, err
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if token == nil {
		err = jwt.ErrSignatureInvalid
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"].(string)
		userId = &sub
		return userId, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
