package CategoryController

import (
	// "golang-restfull-api/model/Category/categorymodel"
	CategoryModel "golang-restfull-api/app/model/Category"
	http "net/http"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func GetUsers(c *gin.Context, db *gorm.DB) {
	var results []CategoryModel.Category
	db.Table("category").Select("id, name, phone_number").Find(&results)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    results,
	})
}
