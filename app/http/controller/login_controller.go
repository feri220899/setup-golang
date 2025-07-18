package controller

import (
	usermodel "golang-restfull-api/app/model/user"
	http "net/http"
	"time"

	gin "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	gorm "gorm.io/gorm"
)

func GetToken(request *gin.Context, db *gorm.DB) {
	username := request.PostForm("username")
	// password := request.PostForm("password")

	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &usermodel.JWTClaim{
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
			Subject:   "access-token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		request.JSON(http.StatusInternalServerError, gin.H{"error": "Token signing failed"})
		return
	}

	request.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"token":    tokenString,
		"expires":  expirationTime,
		"issued":   time.Now(),
		"username": username,
	})
}
