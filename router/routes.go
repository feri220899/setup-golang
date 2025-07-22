package routes

import (
	controller "golang-restfull-api/app/http/controller"
	middleware "golang-restfull-api/app/http/middleware"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	apiGroup := router.Group("/api")
	authorGroup := apiGroup.Group("/auth")
	{
		authorGroup.POST("/login", func(request *gin.Context) {
			controller.GetToken(request, db)
		})

		authorGroup.POST("/refresh-token", func(request *gin.Context) {
			controller.RefreshToken(request, db)
		})
	}
	userGroup := apiGroup.Group("/users")
	{
		userGroup.GET("", middleware.UserMiddleware, func(request *gin.Context) {
			controller.GetUsers(request, db)
		})

		userGroup.POST("/import-data", middleware.UserMiddleware, func(request *gin.Context) {
			controller.ImportData(request, db)
		})

		userGroup.POST("/import-data-progres", middleware.UserMiddleware, func(request *gin.Context) {
			controller.ImportProgres(request, db)
		})
	}

}
