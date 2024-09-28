package routes

import (
	"github.com/beGG1/study_platform_api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	api.GET("/ping", handlers.Ping)
	api.POST("/auth", handlers.Login)
}
