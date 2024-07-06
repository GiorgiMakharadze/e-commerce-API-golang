package routes

import (
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// TO DO - add middlewares

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/logout", auth.Logout)

	}

	return router
}
