package controller

import (
	"errors"
	"fmt"
	"golang-restfull-api/app/helper"
	usermodel "golang-restfull-api/app/model/user"
	http "net/http"
	"strings"
	"time"

	gin "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	gorm "gorm.io/gorm"
)

func generateToken(user usermodel.UserModel, expired int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expired) * time.Minute)
	claims := &usermodel.JWTClaim{
		UserName: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "access_token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(user.Secret_key))
	return tokenString, err
}

func generateRefreshToken(expired int, Secret_key string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expired) * time.Minute)
	rawToken := uuid.New().String()
	combined := rawToken + "|" + expirationTime.Format(time.RFC3339)
	hashedToken, err := helper.Encrypt(combined, Secret_key)
	if err != nil {
		return "", fmt.Errorf("encrypt gagal: %w", err)
	}
	return hashedToken, nil
}

func GetToken(request *gin.Context, db *gorm.DB) {
	username := request.PostForm("username")
	password := request.PostForm("password")
	var user usermodel.UserModel
	db.Table("users").Where("username", username).First(&user)

	if user.Username != username || !helper.CheckPasswordHash(password, user.Password) {
		request.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	} else {
		tokenString, err := generateToken(user, 1)
		if err != nil {
			request.JSON(http.StatusInternalServerError, gin.H{"error": "Token signing failed"})
			return
		}

		refresh_token, err := generateRefreshToken(24*60, user.Secret_key)
		if err != nil {
			request.JSON(http.StatusInternalServerError, gin.H{"error": "Refresh token signing failed"})
			return
		}

		db.Table("users").Where("id", user.Id).Updates(map[string]interface{}{
			"refresh_token": refresh_token,
		})

		request.JSON(http.StatusOK, gin.H{
			"message":       "success",
			"token":         tokenString,
			"refresh_token": refresh_token,
			"issued":        time.Now(),
			"username":      username,
		})
	}
}

func RefreshToken(request *gin.Context, db *gorm.DB) {
	// FROM REQUEST
	request_token := request.GetHeader("refresh_token")
	secret_key := request.GetHeader("x-secret-key")

	// FROM DB
	var user usermodel.UserModel
	result := db.Table("users").Where("refresh_token", request_token).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			request.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
			return
		}
		request.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// CHECK IF EXPIRED
	if user == (usermodel.UserModel{}) {
		request.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// DECRYPT AND CHECK EXPIRATION
	hashed_token, _ := helper.Decrypt(request_token, secret_key)
	parts := strings.SplitN(hashed_token, "|", 2)
	expired_at_part := parts[1]

	expired_at, err := time.Parse(time.RFC3339, expired_at_part)
	if err != nil || time.Now().After(expired_at) {
		request.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
		return
	} else {
		tokenString, err := generateToken(user, 1080)
		if err != nil {
			request.JSON(http.StatusInternalServerError, gin.H{"error": "Token signing failed"})
			return
		} else {
			request.JSON(http.StatusOK, gin.H{
				"message": "success",
				"token":   tokenString,
			})
		}
	}

}
