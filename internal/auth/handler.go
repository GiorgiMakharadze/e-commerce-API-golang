package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUserRole(role string) (UserRole, error) {
	switch role {
	case string(Admin), string(Seller), string(Buyer):
		return UserRole(role), nil
	default:
		return "", errors.New("invalid user role")
	}
}

func RegisterUser(c *gin.Context) {
	var authInput RegisterAuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound User
	if err := db.DB.Where("username = ? OR email = ?", authInput.Username, authInput.Email).First(&userFound).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email already used"})
		return
	}

	if err := db.DB.Where("first_name = ? OR last_name = ?", authInput.FirstName, authInput.LastName).First(&userFound).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "First Name or Last Name already used"})
		return
	}

	role, err := ValidateUserRole(string(authInput.Role))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Choose correct role"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := User{
		Username:  authInput.Username,
		Password:  string(passwordHash),
		Email:     authInput.Email,
		FirstName: authInput.FirstName,
		LastName:  authInput.LastName,
		Role:      role,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func LoginUser(c *gin.Context) {
	var authInput LoginAuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound User
	if err := db.DB.Where("email = ?", authInput.Email).First(&userFound).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	tokenString, err := generateJWT(userFound.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{
		"message":     "Logged in successfully",
		"accessToken": tokenString,
	})
}

func Logout(c *gin.Context) {

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})

}

func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
