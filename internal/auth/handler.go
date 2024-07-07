package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func RegisterUser(c *gin.Context) {
	var authInput RegisterAuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound User
	db.DB.Where("username=?", authInput.Username).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	db.DB.Where("email=?", authInput.Email).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already used"})
		return
	}

	db.DB.Where("first_name=?", authInput.FirstName).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this First Name already used"})
		return
	}

	db.DB.Where("last_name=?", authInput.LastName).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this Last Name already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{
		Username:  authInput.Username,
		Password:  string(passwordHash),
		Email:     authInput.Email,
		FirstName: authInput.FirstName,
		LastName:  authInput.LastName,
		Role:      authInput.Role,
	}

	db.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func LoginUser(c *gin.Context) {
	var authInput LoginAuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound User
	db.DB.Where("email=?", authInput.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	session, err := Store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
	}
	session.Values["user_id"] = userFound.ID
	session.Values["accessToken"] = token

	session.Options.HttpOnly = true
	 session.Save(c.Request, c.Writer)
	jsonSessionValues := make(map[string]interface{})
	for k, v := range session.Values {
		jsonSessionValues[fmt.Sprintf("%v", k)] = v
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Logged in successfully",
		"accessToken": token,
		"session":     session.Values,
	})
}

func Logout(c *gin.Context) {
	session, _ := Store.Get(c.Request, "session")
	userID, ok := session.Values["user_id"].(uint)
	if !ok || userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user is currently logged in"})
		return
	}

	session.Options.MaxAge = -1
	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})

}
