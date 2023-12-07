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
	if err := r.db.Order("name").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
func (r *roleRepository) Create(role *model.Role) (*model.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
func (r *roleRepository) Migrate() error {
	return r.db.AutoMigrate(&model.Role{})
}

func newRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}
