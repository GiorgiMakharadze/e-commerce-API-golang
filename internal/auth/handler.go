package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User Login"})

}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User Logout"})

}
