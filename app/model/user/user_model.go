package user

import "github.com/golang-jwt/jwt/v5"

type UserModel struct {
	UserName     string
	Password     string
	SecretKey    string
	Refreshtoken string
}

type JWTClaim struct {
	UserName string
	jwt.RegisteredClaims
}
