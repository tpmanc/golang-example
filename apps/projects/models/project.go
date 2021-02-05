package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UserId int
	Title string
}

func (model *Project) Validate() bool {
	return false
}
