package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type userRepository struct {
	db *gorm.DB
}

// 获取所有用户
func (u *userRepository) List() (model.Users, error) {
	users := make(model.Users, 0)
	if err := u.db.Omit("Password").Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// 创建用户
func (u *userRepository) Create(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// 更新用户
func (u *userRepository) Update(user *model.User) (*model.User, error) {
	if err := u.db.Model(model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// 更新用户密码
func (u *userRepository) UpdatePassword(user *model.User) error {
	if err := u.db.Model(model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// 删除用户
func (u *userRepository) Delete(user *model.User) error {
	err := u.db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

// 根据ID获取用户
func (u *userRepository) GetUserByID(id uint) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("Password").Preload("Roles").First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// 根据账号获取用户
func (u *userRepository) GetUserByAccount(account string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("Password").Preload("Roles").First(user, "account = ?", account).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// 添加角色
func (u *userRepository) AddRole(user *model.User, role *model.Role) error {
	if err := u.db.Model(user).Association("Roles").Append(role); err != nil {
		return err
	}
	return nil
}

// 删除角色
func (u *userRepository) RemoveRole(user *model.User, role *model.Role) error {
	if err := u.db.Model(user).Association("Roles").Delete(role); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{})
}
func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
