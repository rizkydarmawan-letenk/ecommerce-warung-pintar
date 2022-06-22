package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/api-gateway/routes/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	return router
}
