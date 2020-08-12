package roles

import "github.com/jinzhu/gorm"

type RoleRepository struct {
	DB *gorm.DB
}

func ProvideRoleRepository (DB *gorm.DB) RoleRepository {
	return RoleRepository{DB: DB}
}

func (p *RoleRepository) GetAll() []Role {
	var roles []Role
	p.DB.Find(&roles)

	return roles
}

func (p *RoleRepository) GetById(id uint) Role{
	var role Role
	p.DB.First(&role, id)

	return role
}

func (p *RoleRepository) Save(role Role) Role {
	p.DB.Save(&role)

	return role
}

func (p *RoleRepository) Delete(role Role)  {
	p.DB.Delete(&role)
}

