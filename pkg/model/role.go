package model

type Role struct {
	ID          int64        `gorm:"column:id;primary_key" json:"id"`
	Name        string       `gorm:"column:name;" json:"name"`
	Desc        string       `gorm:"column:desc;" json:"desc"`
	Permissions []Permission `gorm:"many2many:role_permission;" json:"permissions"`
}

func (r *Role) TableName() string {
	return "roles"
}

type Roles []Role

type AddRole struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (a AddRole) GetRole() *Role {
	return &Role{
		Name: a.Name,
		Desc: a.Desc,
	}
}

type Role_Permission struct {
	RoleID       int64 `gorm:"column:role_id;primary_key" json:"role_id"`
	PermissionID int64 `gorm:"column:permission_id;primary_key" json:"permission_id"`
}

func (Role_Permission) TableName() string {
	return "role_permissions"
}
