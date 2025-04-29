package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/darendsen/test-gin/internal/handlers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(db)

	// User routes
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)

	return router
}
