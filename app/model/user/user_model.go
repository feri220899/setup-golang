package user

import (
	excelmodel "golang-restfull-api/app/model/import/excel"

	"github.com/golang-jwt/jwt/v5"
)

type UserModel struct {
	Id            uint                      `gorm:"primaryKey" json:"id"`
	Username      string                    `json:"username"`
	Password      string                    `json:"password"`
	User_key      string                    `json:"user_key"`
	Refresh_token string                    `json:"refresh_token"`
	Import_status []excelmodel.ImportStatus `gorm:"foreignKey:User_id;references:Id" json:"import_status"`
}

type JWTClaim struct {
	UserName string
	User_key string
	jwt.RegisteredClaims
}
