package models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	ServerId int
	Path string
}

func (model *Files) Validate() bool {
	return true
}
