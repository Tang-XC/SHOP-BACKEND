package repository

import (
	"gorm.io/gorm"
)

type repository struct {
	user     UserRepository
	role     RoleRepository
	db       *gorm.DB
	migrants []Migrant
}

func (r *repository) User() UserRepository {
	return r.user
}
func (r *repository) Role() RoleRepository {
	return r.role
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
		db:   db,
		user: newUserRepository(db),
		role: newRoleRepository(db),
	}
	r.migrants = getMigrants(r.user)
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
