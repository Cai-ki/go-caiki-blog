package jwt

import (
	"time"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecretKey = []byte("secret-key")

type Claims struct {
	jwt.RegisteredClaims
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthJwt interface {
	GenerateToken(id uint, username string, email string) (string, error)
	ParseToken(tokenString string) (*Claims, error)
	ValidateClaimsExists(claims *Claims) (bool, error)
}

type authJwtImpl struct {
}

var _ AuthJwt = (*authJwtImpl)(nil)

var Jwt AuthJwt = authJwtImpl{}

func (authJwtImpl) GenerateToken(id uint, username string, email string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-caiki-blog",
		},
		ID:       id,
		Username: username,
		Email:    email,
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

func (authJwtImpl) ValidateClaimsExists(claims *Claims) (bool, error) {
	db := storage.DB.GetDB()
	var user models.Users
	if err := db.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	}

	return true, nil
}
