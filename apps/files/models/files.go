package models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	ServerId int
	Path string
}

func (m *Files) Validate() (bool, string) {
	if m.ServerId <= 0 {
		return false, "serverId is required"
	}
	if len(m.Path) < 2 {
		return false, "path is required"
	}

	return true, ""
}
