package service

import (
	"shop/pkg/model"
	"shop/pkg/repository"
)

type roleService struct {
	roleRepository repository.RoleRepository
}

func (r *roleService) List() (model.Roles, error) {
	roles, err := r.roleRepository.List()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{roleRepository: roleRepository}
}
