package service

import (
	"mime/multipart"
	"shop/pkg/model"
)

type UserService interface {
	List() (model.Users, error)
	GetUser(string) (*model.User, error)
	GetUserById(string) (*model.User, error)
	Create(*model.User) (string, error)
	Update(string, *model.User) (*model.User, error)
	Delete(string) error
	UpdatePassword(string, *model.UpdatedPassword) (string, error)
	AddRole(uint, uint) error
	RemoveRole(uint, uint) error
}

type AuthService interface {
	Login(*model.LoginUser) (interface{}, error)
}

type RoleService interface {
	List() (model.Roles, error)
	Create(*model.Role) (string, error)
}

type PermissionService interface {
	List() (model.Permissions, error)
	Create(*model.Permission) (string, error)
}
type ProductService interface {
	List(page int, size int, category int) (model.ProductsResponse, error)
	Create(product *model.AddProduct, user *model.User) (string, error)
	Update(uint, *model.Product) (*model.Product, error)
	Delete(uint) error
	GetProductByID(uint) (*model.Product, error)
	AddFile(uint, uint) error
}
type CategoryService interface {
	List() (model.Categorys, error)
	Create(*model.Category) (string, error)
	Delete(uint) error
}
type UploadService interface {
	UploadImages(files []*multipart.FileHeader, token string) ([]model.FileResponse, error)
}
type FileService interface {
	List() (model.Files, error)
	Create(files []*multipart.FileHeader, token string) (model.Files, error)
	Delete(file *model.File) error
	GetFileByID(uint) (*model.File, error)
}
