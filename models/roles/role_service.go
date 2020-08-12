package roles

type RoleService struct {
	RoleRepository RoleRepository
}

func ProvideRoleService(repository RoleRepository) RoleService {
	return RoleService{RoleRepository: repository}
}

func (p *RoleService) GetAll() []Role{
	return p.RoleRepository.GetAll()
}

func (p *RoleService) GetById(id uint) Role{
	return p.RoleRepository.GetById(id)
}

func (p *RoleService) Save(role Role) Role{
	return p.RoleRepository.Save(role)
}

func (p *RoleService) Delete(role Role)  {
	p.RoleRepository.Delete(role)
}