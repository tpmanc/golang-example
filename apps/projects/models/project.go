package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UserId int
	Title string
}

func (m *Project) Validate() (bool, string) {
	if m.UserId <= 0 {
		return false, "userId is required"
	}

	if len(m.Title) < 2 {
		return false, "title is required"
	}

	return true, ""
}
