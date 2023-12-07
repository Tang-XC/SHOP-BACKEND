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
func (r *roleService) Create(role *model.Role) (string, error) {
	message := "创建成功"
	_, err := r.roleRepository.Create(role)
	if err != nil {
		return "", err
	}
	return message, nil
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{roleRepository: roleRepository}
}
