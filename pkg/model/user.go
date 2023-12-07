package model

import (
	"encoding/json"
)

type User struct {
	ID       uint   `gorm:"column:id;" json:"id"`
	Name     string `gorm:"column:name;" json:"name"`
	Account  string `gorm:"column:account;" json:"account"'`
	Phone    string `gorm:"column:phone;" json:"phone"'`
	Desc     string `gorm:"column:desc;" json:"desc"`
	PassWord string `gorm:"column:password;" json:"password"`
	Email    string `gorm:"column:email;" json:"email"`
	Avatar   string `gorm:"column:avatar;" json:"avatar"`
	Region   string `gorm:"column:region;" json:"region"`
	Address  string `gorm:"column:address;" json:"address"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
}

func (table User) TableName() string {
	return "users"
}

type Users []User

type LoginUser struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (c *RegisterUser) GetUser() *User {
	return &User{
		Account:  c.Account,
		PassWord: c.Password,
		Phone:    c.Phone,
		Email:    c.Email,
	}
}

type UpdatedUser struct {
	ID      uint     `json:"id"`
	Name    string   `json:"name"`
	Account string   `json:"account"'`
	Desc    string   `json:"desc"`
	Phone   string   `json:"phone"'`
	Email   string   `json:"email"`
	Avatar  string   `json:"avatar"`
	Region  []string `json:"region"`
	Address string   `json:"address"`
}

func (c *UpdatedUser) GetUser() *User {
	region, err := json.Marshal(c.Region)
	if err != nil {
		return nil
	}
	return &User{
		ID:      c.ID,
		Account: c.Account,
		Name:    c.Name,
		Desc:    c.Desc,
		Avatar:  c.Avatar,
		Email:   c.Email,
		Phone:   c.Phone,
		Region:  string(region),
		Address: c.Address,
	}
}

type UpdatedPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
