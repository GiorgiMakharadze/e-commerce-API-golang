package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func CSRFMiddleware(secret string) gin.HandlerFunc {
	csrfMiddleware := csrf.Protect(
		[]byte(secret),
		csrf.Secure(false),
	)
	return func(c *gin.Context) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("x-CSRF-Token", csrf.Token(r))
			for k, v := range w.Header() {
				c.Writer.Header().Set(k, v[0])
			}
			c.Request = r
			c.Next()
		})
		csrfMiddleware(h).ServeHTTP(c.Writer, c.Request)
	}
}
