package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type permissionRepository struct {
	db *gorm.DB
}

func (u *permissionRepository) List() (model.Permissions, error) {
	permissions := make(model.Permissions, 0)
	if er := u.db.Order("name").Find(&permissions).Error; er != nil {
		return nil, er
	}
	return permissions, nil
}
func (u *permissionRepository) Create(permission *model.Permission) (*model.Permission, error) {
	if err := u.db.Create(permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}
func (u *permissionRepository) Migrate() error {
	return u.db.AutoMigrate(&model.Permission{})
}
func newPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}
