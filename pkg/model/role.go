package model

type Role struct {
	ID   int64
	Name string
}

func (r *Role) TableName() string {
	return "roles"
}

type Roles []Role
