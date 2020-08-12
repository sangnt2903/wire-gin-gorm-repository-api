package main

import (
	"github.com/jinzhu/gorm"
	"golang-repository-gin-gorm-wire/models/roles"
	"github.com/google/wire"
)

func initRoleApi(db *gorm.DB) roles.RoleApi{
	wire.Build(roles.ProvideRoleApi, roles.ProvideRoleRepository, roles.ProvideRoleService)

	return roles.RoleApi{}
}
