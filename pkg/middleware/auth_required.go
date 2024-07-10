package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func AuthRequired(c *gin.Context) {
	session, _ := Store.Get(c.Request, "auth-session")

	if session.Values["accessToken"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
