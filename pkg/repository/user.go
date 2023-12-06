package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) List() (model.Users, error) {
	users := make(model.Users, 0)
	//if err := u.db.Preload(model.UserAuthInfoAssociation).Preload(model.GroupAssociation).Preload("Roles").Order("name").Find(&users).Error; err != nil {
	//	return nil, err
	//}
	if err := u.db.Order("name").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepository) Create(user *model.User) (*model.User, error) {
	var userCreateField []string = []string{"name", "email", "password", "avatar", "account", "phone"}
	if err := u.db.Select(userCreateField).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userRepository) Update(user *model.User) (*model.User, error) {
	if err := u.db.Model(model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userRepository) UpdatePassword(user *model.User) error {
	if err := u.db.Model(model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepository) Delete(user *model.User) error {
	err := u.db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *userRepository) GetUserByID(id uint) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("Password").Preload("Roles").First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userRepository) GetUserByAccount(account string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("Password").Where("account = ?", account).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//	func (u *userRepository) AddRole(role *model.Role, user *model.User) error {
//		return u.db.Model(user).Association("Roles").Append(role)
//	}
//
//	func (u *userRepository) DelRole(role *model.Role, user *model.User) error {
//		return u.db.Model(user).Association("Roles").Delete(role)
//	}
func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{})
}
func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
