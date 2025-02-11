package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecretKey = []byte("secret-key")

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

type AuthJwt interface {
	GenerateToken(username string) (string, error)
	ParseToken(tokenString string) (*Claims, error)
}

type authJwtImpl struct {
}

var _ AuthJwt = (*authJwtImpl)(nil)

var Jwt AuthJwt = authJwtImpl{}

func (authJwtImpl) GenerateToken(username string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-caiki-blog",
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}

func (authJwtImpl) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
