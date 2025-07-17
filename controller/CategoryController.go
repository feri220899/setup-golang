package controller

import (
	// "golang-restfull-api/model/Category/categorymodel"
	CategoryModel "golang-restfull-api/model/Category"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context, db *gorm.DB) {
	var results []CategoryModel.Category
	db.Table("category").Select("id, name, phone_number").Find(&results)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    results,
	})
}
