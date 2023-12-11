package service

import (
	"errors"
	"gorm.io/gorm"
	"shop/pkg/model"
	"shop/pkg/repository"
	utils "shop/pkg/utils/token"
	"strconv"
)

type userService struct {
	userRepository repository.UserRepository
	roleRepository repository.RoleRepository
}

func (u *userService) List() (model.Users, error) {
	return u.userRepository.List()
}
func (u *userService) GetUser(token string) (*model.User, error) {
	claim, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return u.userRepository.GetUserByAccount(claim.Subject)
}
func (u *userService) Create(user *model.User) (string, error) {
	_, err := u.userRepository.GetUserByAccount(user.Account)
	message := "用户已存在"
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user, err := u.userRepository.Create(user)
			if err != nil {
				return "", err
			}
			message = "用户创建成功"
			//创建用户成功后，给用户添加默认角色
			if err := u.AddRole(user.ID, 2); err != nil {
				message = "用户创建成功，但是添加默认角色失败"
			}

			return message, nil
		}
	}
	return "", errors.New(message)
}

func (u *userService) Update(id string, new *model.User) (*model.User, error) {
	old, err := u.GetUserById(id)
	if err != nil {
		return nil, err
	}
	new.ID = old.ID
	return u.userRepository.Update(new)
}
func (u *userService) Delete(id string) error {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	user := &model.User{
		ID: uint(userId),
	}
	return u.userRepository.Delete(user)
}
func (u *userService) UpdatePassword(id string, up *model.UpdatedPassword) (string, error) {
	var user = new(model.User)
	old, err := u.GetUserById(id)
	if err != nil {
		return "", err
	}
	if old.PassWord != up.OldPassword {
		return "", errors.New("旧密码错误")
	}
	user.ID = old.ID
	if up.NewPassword == up.OldPassword {
		return "", errors.New("新旧密码不能相同")
	}
	user.PassWord = up.NewPassword
	err = u.userRepository.UpdatePassword(user)
	if err != nil {
		return "", err
	}
	return "更改成功", nil
}
func (u *userService) GetUserById(id string) (*model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return u.userRepository.GetUserByID(uint(uid))
}
func (u *userService) AddRole(userId uint, roleId uint) error {
	//查找用户
	user, err := u.userRepository.GetUserByID(userId)
	if err != nil {
		return errors.New("用户不存在")
	}
	//查找角色
	role, err := u.roleRepository.GetRoleByID(roleId)
	if err != nil {
		return errors.New("角色不存在")
	}
	//添加角色
	err = u.userRepository.AddRole(user, role)
	if err != nil {
		return err
	}
	return nil
}
func (u *userService) RemoveRole(userId uint, roleId uint) error {
	//查找用户
	user, err := u.userRepository.GetUserByID(userId)
	if err != nil {
		return errors.New("用户不存在")
	}
	//查找角色
	role, err := u.roleRepository.GetRoleByID(roleId)
	if err != nil {
		return errors.New("角色不存在")
	}
	//删除角色
	err = u.userRepository.RemoveRole(user, role)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userRepository repository.UserRepository, roleRepository repository.RoleRepository) UserService {
	return &userService{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}
