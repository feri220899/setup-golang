package routes

import (
	CategoryController "golang-restfull-api/app/http/controller"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("", func(c *gin.Context) {
			CategoryController.GetUsers(c, db)
		})
		// Tambah PUT, DELETE, GET by ID
	}
}
