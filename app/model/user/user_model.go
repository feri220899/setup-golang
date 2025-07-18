package user

import "github.com/golang-jwt/jwt/v5"

type UserModel struct {
	Id            uint `gorm:"primaryKey"`
	Username      string
	Password      string
	Secret_key    string
	Refresh_token string
}

type JWTClaim struct {
	UserName string
	jwt.RegisteredClaims
}
