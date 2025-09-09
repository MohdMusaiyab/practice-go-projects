package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// CustomClaims defines the structure of JWT payload
type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new token for given user
func GenerateJWT(userID uint, email string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // 3 days expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "blog-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT checks if token string is valid
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
}

// ParseToken parses a token and extracts claims
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := ValidateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
