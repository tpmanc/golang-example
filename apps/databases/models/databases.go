package models

import "gorm.io/gorm"

type Databases struct {
	gorm.Model
	ServerId int
	User string
	Password string
	Database string
}

func (m *Databases) Validate() (bool, string) {
	if m.ServerId <= 0 {
		return false, "serverID is required"
	}
	if len(m.User) < 2 {
		return false, "user is required"
	}
	if len(m.Password) < 2 {
		return false, "password is required"
	}
	if len(m.Database) < 2 {
		return false, "database name is required"
	}

	return true, ""
}
