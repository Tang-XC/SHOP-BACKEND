package repository

import (
	"shop/pkg/model"
)

type Repository interface {
	User() UserRepository
	Role() RoleRepository
	Permission() PermissionRepository
	Category() CategoryRepository
	Product() ProductRepository
	File() FileRepository
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
	AddRole(*model.User, *model.Role) error
	RemoveRole(*model.User, *model.Role) error
	Migrate() error
}

type RoleRepository interface {
	List() (model.Roles, error)
	Create(*model.Role) (*model.Role, error)
	GetRoleByID(uint) (*model.Role, error)
	Migrate() error
}
type PermissionRepository interface {
	List() (model.Permissions, error)
	Create(*model.Permission) (*model.Permission, error)
	Migrate() error
}
type CategoryRepository interface {
	List() (model.Categorys, error)
	Create(*model.Category) (*model.Category, error)
	Delete(id uint) error
	GetCategoryByID(uint) (*model.Category, error)
	Migrate() error
}
type ProductRepository interface {
	List(page int, size int, category int, keywords string) (model.ProductsResponse, error)
	Create(*model.Product) (*model.Product, error)
	Update(*model.Product) (*model.Product, error)
	Delete(*model.Product) error
	RemoveFileRelation(*model.Product) error
	GetProductByID(uint) (*model.Product, error)
	AddFile(*model.Product, *model.File) error
	Migrate() error
}
type FileRepository interface {
	List() (model.Files, error)
	Create(*model.File) (*model.File, error)
	Delete(*model.File) error
	GetFileByID(uint) (*model.File, error)
	Migrate() error
}
