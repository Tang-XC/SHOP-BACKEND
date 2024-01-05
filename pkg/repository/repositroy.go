package repository

import (
	"gorm.io/gorm"
)

type repository struct {
	user       UserRepository
	role       RoleRepository
	permission PermissionRepository
	category   CategoryRepository
	product    ProductRepository
	file       FileRepository
	db         *gorm.DB
	migrants   []Migrant
}

func (r *repository) User() UserRepository {
	return r.user
}
func (r *repository) Role() RoleRepository {
	return r.role
}
func (r *repository) Permission() PermissionRepository {
	return r.permission
}
func (r *repository) Category() CategoryRepository {
	return r.category
}
func (r *repository) Product() ProductRepository {
	return r.product
}
func (r *repository) File() FileRepository {
	return r.file
}
func (r *repository) Init() error {
	return nil
}
func (r *repository) Close() error {
	db, _ := r.db.DB()
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
func (r *repository) Migrate() error {
	for _, m := range r.migrants {
		if err := m.Migrate(); err != nil {
			return err
		}
	}
	return nil
}

func NewRepository(db *gorm.DB) Repository {
	r := &repository{
		db:         db,
		user:       newUserRepository(db),
		role:       newRoleRepository(db),
		permission: newPermissionRepository(db),
		category:   newCategoryRepository(db),
		product:    newProductRepository(db),
		file:       newFileRepository(db),
	}
	r.migrants = getMigrants(r.user, r.permission, r.role, r.product, r.category, r.file)
	return r
}

// 获取所有的迁移对象
func getMigrants(objs ...interface{}) []Migrant {
	var migrants []Migrant
	for _, obj := range objs {
		if m, ok := obj.(Migrant); ok {
			migrants = append(migrants, m)
		}
	}
	return migrants
}
