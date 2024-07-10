package handler

import (
	"log"
	"net/http"
	"os"

	auth_service "github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/service"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
	sessions_service "github.com/GiorgiMakharadze/e-commerce-API-golang/internal/sessions/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

type AuthHandler struct {
	authService    auth_service.AuthService
	sessionService sessions_service.SessionService
}

func NewAuthHandler(authService auth_service.AuthService, sessionService sessions_service.SessionService) *AuthHandler {
	return &AuthHandler{
		authService:    authService,
		sessionService: sessionService,
	}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var authInput models.RegisterAuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.RegisterUser(authInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	log.Println("Session Key:", os.Getenv("SESSION_KEY"))

	var authInput models.LoginAuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.authService.LoginUser(authInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.GetUserByEmail(authInput.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	session := &models.Session{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	h.sessionService.CreateSession(session)

	gSession, _ := Store.Get(c.Request, "auth-session")
	gSession.Values["accessToken"] = accessToken
	gSession.Values["refreshToken"] = refreshToken
	gSession.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{
		"message":      "Logged in successfully",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	log.Println("Session Key:", os.Getenv("SESSION_KEY"))

	gSession, err := Store.Get(c.Request, "auth-session")
	if err != nil {
		log.Println("Error retrieving session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}

	accessToken, ok := gSession.Values["accessToken"].(string)
	if !ok || accessToken == "" {
		log.Println("Invalid session token")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session"})
		return
	}

	err = h.sessionService.DeleteSessionByToken(accessToken)
	if err != nil {
		log.Println("Error deleting session from database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	gSession.Options.MaxAge = -1
	if err := gSession.Save(c.Request, c.Writer); err != nil {
		log.Println("Error saving session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}
