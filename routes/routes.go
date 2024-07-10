package routes

import (
	"github.com/GiorgiMakharadze/e-commerce-API-golang/config"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/db"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/handler"
	authRepo "github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/repository"
	authService "github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/service"
	sessionRepo "github.com/GiorgiMakharadze/e-commerce-API-golang/internal/sessions/repository"
	sessionService "github.com/GiorgiMakharadze/e-commerce-API-golang/internal/sessions/service"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authRepo := authRepo.NewAuthRepository(db.DB)
	sessionsRepo := sessionRepo.NewSessionRepository(db.DB)
	authService := authService.NewAuthService(authRepo)
	sessionService := sessionService.NewSessionService(sessionsRepo)
	authHandler := handler.NewAuthHandler(authService, sessionService)

	router.Use(middleware.CSRFMiddleware(config.AppConfig.Csrf_key))

	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.GET("/csrf-token", authHandler.GetCSRFToken)
		authRoutes.POST("/register", authHandler.RegisterUser)
		authRoutes.POST("/login", authHandler.LoginUser)
		authRoutes.POST("/logout", middleware.AuthRequired, authHandler.Logout)

	}

	protectedProductRoute := router.Group("/api/v1/product")
	protectedProductRoute.Use(middleware.AuthRequired)
	{
		protectedProductRoute.POST("/create-product")
	}

	return router
}
