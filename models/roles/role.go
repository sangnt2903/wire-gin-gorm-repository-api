package roles

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	RoleName	string `json:"role_name"`
}
