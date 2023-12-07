package model

type Role struct {
	ID          int64        `gorm:"column:id;primary_key" json:"id"`
	Name        string       `gorm:"column:name;" json:"name"`
	Desc        string       `gorm:"column:desc;" json:"desc"`
	Key         uint8        `gorm:"column:key;" json:"key"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type AddRole struct {
	Name          string  `json:"name"`
	Desc          string  `json:"desc"`
	Key           uint8   `json:"key"`
	PermissionIds []int64 `json:"permission_ids"`
}

func (a AddRole) GetRole() *Role {
	return &Role{
		Name: a.Name,
		Key:  a.Key,
		Desc: a.Desc,
	}
}

func (r *Role) TableName() string {
	return "roles"
}

type Roles []Role
