package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTItf interface {
	CreateToken(UserId uuid.UUID, IsAdmin bool) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, bool, error)
}

type JWTService struct {
	SecretKey string
	Expired   time.Duration
}

type Claims struct {
	UserId  uuid.UUID
	IsAdmin bool
	jwt.RegisteredClaims
}

func NewJWT(secretKey string, expiredTime int) JWTItf {
	return &JWTService{
		SecretKey: secretKey,
		Expired:   time.Duration(expiredTime) * time.Hour,
	}
}

func (w *JWTService) CreateToken(UserId uuid.UUID, IsAdmin bool) (string, error) {
	claims := &Claims{
		UserId:  UserId,
		IsAdmin: IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(w.Expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (w *JWTService) ValidateToken(tokenString string) (uuid.UUID, bool, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(w.SecretKey), nil
	})

	if err != nil {
		return claims.UserId, false, err
	}

	if !token.Valid {
		return claims.UserId, false, err
	}

	return claims.UserId, claims.IsAdmin, nil
}
