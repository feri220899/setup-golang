package routes

import (
	"golang-restfull-api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("", func(c *gin.Context) {
			controller.GetUsers(c, db)
		})
		// Tambah PUT, DELETE, GET by ID
	}
}
