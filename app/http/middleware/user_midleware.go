package usermiddleware

import (
	"fmt"
	"net/http"
	"strings"

	usermodel "golang-restfull-api/app/model/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func UserMiddleware(request *gin.Context, db *gorm.DB) {
	auth_header := request.GetHeader("X-token")
	secret_key := viper.GetString("API_SECRET_KEY")
	user_key := request.GetHeader("user_key")

	if auth_header == "" || !strings.HasPrefix(auth_header, "") || user_key == "" {
		request.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User key or token required"})
		return
	}

	// FROM DB
	var user usermodel.UserModel
	user_key_db := db.Table("users").Where("user_key", user_key).First(&user)
	if user_key_db.Error != nil {
		request.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User key not found"})
		return
	}

	tokenStr := strings.TrimPrefix(auth_header, "")
	claims, err := ValidateJWT(tokenStr, secret_key)
	if err != nil || claims.User_key != user.User_key || claims.UserName != user.Username {
		request.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}
	request.Next()
}

func ValidateJWT(tokenStr string, secret_key string) (*usermodel.JWTClaim, error) {
	claims := &usermodel.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret_key), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
