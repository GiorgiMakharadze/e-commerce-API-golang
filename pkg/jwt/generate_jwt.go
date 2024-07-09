package jwt

import (
	"os"
	"time"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID uint, role models.UserRole) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"id":   userID,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}
