package routes

import (
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.POST("/register", auth.RegisterUser)
		authRoutes.POST("/login", auth.LoginUser)
		authRoutes.POST("/logout", auth.Logout)

	}

	protectedProductRoute := router.Group("/api/v1/product")
	protectedProductRoute.Use(middleware.AuthRequired())
	{
		protectedProductRoute.POST("/create-product")
	}

	return router
}
