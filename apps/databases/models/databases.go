package models

import "gorm.io/gorm"

type Databases struct {
	gorm.Model
	ServerId int
	User string
	Password string
	Database string
}

func (model *Databases) Validate() bool {
	return true
}
