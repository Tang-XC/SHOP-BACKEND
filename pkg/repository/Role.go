package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) List() (model.Roles, error) {
	roles := make(model.Roles, 0)
	if err := r.db.Order("name").Find(roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func newRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}
