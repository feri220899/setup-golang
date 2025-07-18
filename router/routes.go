package routes

import (
	controller "golang-restfull-api/app/http/controller"
	middleware "golang-restfull-api/app/http/middleware"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("", middleware.UserMiddleware, func(request *gin.Context) {
			controller.GetUsers(request, db)
		})
	}

	authorGroup := router.Group("/auth")
	{
		authorGroup.POST("/login", func(request *gin.Context) {
			controller.GetToken(request, db)
		})

		authorGroup.POST("/refresh-token", func(request *gin.Context) {
			controller.RefreshToken(request, db)
		})
	}

}
