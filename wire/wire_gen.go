
package wire

import (
	"github.com/jinzhu/gorm"
	"golang-repository-gin-gorm-wire/models/roles"
)

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Injectors from wire.go:
func InitRoleApi(db *gorm.DB) roles.RoleApi {
	roleRepository := roles.ProvideRoleRepository(db)
	roleService := roles.ProvideRoleService(roleRepository)
	roleApi := roles.ProvideRoleApi(roleService)
	return roleApi
}