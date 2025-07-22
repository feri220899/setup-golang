package routes

import (
	controller "golang-restfull-api/app/http/controller"
	middleware "golang-restfull-api/app/http/middleware"
	config "golang-restfull-api/config"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	// Create route groups
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", config.Routes(db, controller.GetToken))
			auth.POST("/refresh-token", config.Routes(db, controller.RefreshToken))
		}
		users := api.Group("/users").Use(config.Routes(db, middleware.UserMiddleware))
		{
			users.GET("", config.Routes(db, controller.GetUsers))
			users.POST("/import-data", config.Routes(db, controller.ImportData))
			users.POST("/import-data-progress", config.Routes(db, controller.ImportProgres))
		}
	}
}
