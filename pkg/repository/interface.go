package repository

import (
	"shop/pkg/model"
)

type Repository interface {
	User() UserRepository
	Role() RoleRepository
	Permission() PermissionRepository
	Init() error
	Close() error
	Migrant
}
type Migrant interface {
	Migrate() error
}

type UserRepository interface {
	GetUserByID(uint) (*model.User, error)
	GetUserByAccount(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	UpdatePassword(*model.User) error
	Delete(*model.User) error
	Migrate() error
}

type RoleRepository interface {
	List() (model.Roles, error)
	Create(*model.Role) (*model.Role, error)
	Migrate() error
}
type PermissionRepository interface {
	List() (model.Permissions, error)
	Create(*model.Permission) (*model.Permission, error)
	Migrate() error
}
