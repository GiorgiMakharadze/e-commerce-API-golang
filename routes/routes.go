package routes

import (
	"github.com/GiorgiMakharadze/e-commerce-API-golang/db"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/handler"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/repository"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/service"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authRepo := repository.NewAuthRepository(db.DB)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.POST("/register", authHandler.RegisterUser)
		authRoutes.POST("/login", authHandler.LoginUser)
		authRoutes.POST("/logout", authHandler.Logout)

	}

	protectedProductRoute := router.Group("/api/v1/product")
	protectedProductRoute.Use(middleware.AuthRequired())
	{
		protectedProductRoute.POST("/create-product")
	}

	return router
}
