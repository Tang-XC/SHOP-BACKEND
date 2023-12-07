package service

import "shop/pkg/model"

type UserService interface {
	List() (model.Users, error)
	GetUser(string) (*model.User, error)
	GetUserById(string) (*model.User, error)
	Create(*model.User) (string, error)
	Update(string, *model.User) (*model.User, error)
	Delete(string) error
	UpdatePassword(string, *model.UpdatedPassword) (string, error)
	//AddRole(id, rid string) error
	//DelRole(id, rid string) error
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
