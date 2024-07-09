package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func AuthRequired(c *gin.Context) {
	session, _ := store.Get(c.Request, "auth-session")

	if session.Values["accessToken"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
