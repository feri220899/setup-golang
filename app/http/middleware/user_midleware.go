package usermiddleware

import (
	"fmt"
	"net/http"
	"strings"

	usermodel "golang-restfull-api/app/model/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserMiddleware(request *gin.Context) {
	authHeader := request.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		request.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	_, err := ValidateJWT(tokenStr)
	if err != nil {
		request.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}
	request.Next()
}

func ValidateJWT(tokenStr string) (*usermodel.JWTClaim, error) {
	claims := &usermodel.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("your-secret-key"), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
