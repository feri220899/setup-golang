package routes

import (
	authcontroller "golang-restfull-api/app/http/controller/auth"
	categorycontroller "golang-restfull-api/app/http/controller/category"
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
			auth.POST("/login", config.Routes(db, authcontroller.GetToken))
			auth.POST("/refresh-token", config.Routes(db, authcontroller.RefreshToken))
		}
		users := api.Group("/users").Use(config.Routes(db, middleware.UserMiddleware))
		{
			users.GET("", config.Routes(db, categorycontroller.GetUsers))
			users.POST("/import-data", config.Routes(db, categorycontroller.ImportData))
			users.POST("/import-data-progress", config.Routes(db, categorycontroller.ImportProgres))
			users.GET("/get-data-import/:id", config.Routes(db, categorycontroller.GetDataImport))
		}
	}
}
