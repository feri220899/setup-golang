package user

import (
	excelmodel "golang-restfull-api/app/model/import/excel"

	"github.com/golang-jwt/jwt/v5"
)

type UserModel struct {
	Id            uint `gorm:"primaryKey"`
	Username      string
	Password      string
	User_key      string
	Refresh_token string
	Import_status []excelmodel.ImportStatus `gorm:"foreignKey:User_id;references:Id"`
}

type JWTClaim struct {
	UserName string
	jwt.RegisteredClaims
}
