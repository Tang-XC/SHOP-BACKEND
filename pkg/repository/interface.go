package repository

import (
	"shop/pkg/model"
)

type Repository interface {
	User() UserRepository
	Role() RoleRepository
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
	//AddRole(role *model.Role, user *model.User) error
	//DelRole(role *model.Role, user *model.User) error
	Migrate() error
}

type RoleRepository interface {
	List() (model.Roles, error)
}
