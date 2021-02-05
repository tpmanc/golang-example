package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	ProjectId int
	Host string
	User string
	Password string
	Port string
}

func (model *Server) Validate() bool {
	return true
}
