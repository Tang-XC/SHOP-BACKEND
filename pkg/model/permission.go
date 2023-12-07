package model

// 定义权限结构体
type Permission struct {
	ID    int64  `gorm:"column:id;primary_key" json:"id"`
	Name  string `gorm:"column:name;" json:"name"`
	Tag   string `gorm:"column:tag;" json:"tag"`
	Desc  string `gorm:"column:desc;" json:"desc"`
	Roles []Role `gorm:"many2many:role_permissions;"`
}

func (p *Permission) TableName() string {
	return "permissions"
}

type Permissions []Permission

type AddPermission struct {
	Name string `json:"name" `
	Tag  string `json:"tag" `
	Desc string `json:"desc"`
}

func (a *AddPermission) GetPermission() *Permission {
	return &Permission{
		Name: a.Name,
		Tag:  a.Tag,
		Desc: a.Desc,
	}
}
