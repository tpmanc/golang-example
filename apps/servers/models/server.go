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

func (m *Server) Validate() (bool, string) {
	if m.ProjectId <= 0 {
		return false, "projectId is required"
	}

	if len(m.Host) < 2 {
		return false, "host is required"
	}

	if len(m.User) < 2 {
		return false, "user is required"
	}

	if len(m.Password) < 2 {
		return false, "password is required"
	}

	if len(m.Port) < 2 {
		return false, "port is required"
	}

	return true, ""
}
