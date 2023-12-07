package service

import (
	"shop/pkg/model"
	"shop/pkg/repository"
)

type permissionService struct {
	permissionRepository repository.PermissionRepository
}

func (p *permissionService) List() (model.Permissions, error) {
	return p.permissionRepository.List()
}
func (p *permissionService) Create(permission *model.Permission) (string, error) {
	message := "创建成功"
	_, err := p.permissionRepository.Create(permission)
	if err != nil {
		return "", err
	}
	return message, nil
}

func NewPermissionService(permissionRepository repository.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepository: permissionRepository,
	}
}
