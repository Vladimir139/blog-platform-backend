package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Секреты для токенов берём из переменных окружения
var (
	accessTokenSecret  = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	refreshTokenSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
)

// GenerateAccessToken генерирует короткоживущий JWT access токен.
// В поле Subject сохраняется ID пользователя, а не email.
func GenerateAccessToken(userID string) (string, error) {
	expirationTime := time.Now().Add(time.Hour) // 1 час жизни access токена
	claims := &jwt.StandardClaims{
		Subject:   userID, // Используем ID пользователя
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "blog-platform",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessTokenSecret)
}

// GenerateRefreshToken генерирует долгоживущий JWT refresh токен.
func GenerateRefreshToken(userID string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 7 дней
	claims := &jwt.StandardClaims{
		Subject:   userID, // Используем ID пользователя
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "blog-platform",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshTokenSecret)
}

// ParseRefreshToken парсит и валидирует refresh токен
func ParseRefreshToken(tokenStr string) (*jwt.StandardClaims, error) {
	return parseToken(tokenStr, refreshTokenSecret)
}

// ParseAccessToken парсит и валидирует access токен, если понадобится
func ParseAccessToken(tokenStr string) (*jwt.StandardClaims, error) {
	return parseToken(tokenStr, accessTokenSecret)
}

func parseToken(tokenStr string, secret []byte) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
