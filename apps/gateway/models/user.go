package models

import (
	"gorm.io/gorm"
)

type UserApi struct {
	Username string `json:"username"`
}
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
