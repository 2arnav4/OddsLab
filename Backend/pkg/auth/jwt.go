package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
Define Claims Struct:
Define a custom struct Claims that embeds jwt.RegisteredClaims.
Add fields for UserID (string) and Email (string) with custom JSON tags.
*/

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

/*
Generate Token Function (GenerateToken):
Define GenerateToken(UserID string, email string) (string, error).
Read the JWT_SECRET from environment variables.
Create a Claims instance, setting the UserID, Email, and the standard claims (like expiration time ExpiresAt set to 24 hours, and issue time IssuedAt).
Create a new token using jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
Sign the token using token.SignedString([]byte(secret)).
*/

func GenerateToken(UserID string, email string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}
	claims := Claims{
		UserID: UserID,
		Email:  email,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

/*
Validate Token Function (ValidateToken):
Define ValidateToken(tokenStr string) (*Claims, error).
Read the JWT_SECRET.
Parse the token string using jwt.ParseWithClaims, passing a callback function that verifies the signing method is HMAC and returns the secret key.
Check if the token is valid and extract the claims. Return the claims or an error.
*/

func ValidateToken(tokenStr string) (*Claims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable is not set")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(jwtSecret), nil

	})
	if err != nil {
		return nil, err

	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid claims")

	}
	return claims, nil
}
