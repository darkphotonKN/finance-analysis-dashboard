package auth

import (
	"time"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/config"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, string, error) {
	// Set custom and standard claims
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Access token expires in 24 hours
		},
	}

	refreshClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
		},
	}

	// Create the token using secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Generate the signed token string
	tokenString, err := token.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

func ParseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecretKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
