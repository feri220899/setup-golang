package asersicontroller

import (
	"encoding/base64"
	"fmt"
	"golang-restfull-api/app/helper"
	asersimodel "golang-restfull-api/app/model/asersi"
	usermodel "golang-restfull-api/app/model/user"
	"golang-restfull-api/app/service"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ImportData(request *gin.Context, db *gorm.DB) {
	file, err := request.FormFile("file")
	if err != nil {
		helper.AbortWithJSON(request, 400, "error", "File tidak ditemukan", nil)
		return
	}
	user_key := request.GetHeader("user_key")
	var user usermodel.UserModel
	user_key_db := db.Table("asersi.users").Where("user_key", user_key).First(&user)
	if user_key_db.Error != nil {
		request.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User key not found"})
		return
	}

	file_name := filepath.Ext(file.Filename)
	path := helper.PublicPath("uploads/asersi", fmt.Sprintf("data_%d%s", time.Now().Unix(), file_name))
	request.SaveUploadedFile(file, path)
	path_encode := base64.StdEncoding.EncodeToString([]byte(path))

	totalRows := service.CekTemplate(path_encode, "asersi_cvc", request)
	if totalRows == 0 {
		path_decode, _ := base64.StdEncoding.DecodeString(path_encode)
		os.Remove(string(path_decode))
		return
	}

	import_data := asersimodel.ImportStatus{
		Import_file_path: path_encode,
		Import_status:    "processing",
		Import_start:     1,
		Import_batch:     1000,
		Import_total_row: totalRows,
		User_id:          user.Id,
	}

	if err := db.Table("asersi.import_status").Create(&import_data).Error; err != nil {
		helper.AbortWithJSON(request, 500, "error", "Failed to save import status", nil)
		path_decode, _ := base64.StdEncoding.DecodeString(path_encode)
		os.Remove(string(path_decode))
		return
	}

	request.JSON(http.StatusOK, gin.H{
		"message":          "File uploaded successfully",
		"import_status_id": import_data.Id,
	})

}
