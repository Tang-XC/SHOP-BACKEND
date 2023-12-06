package service

import (
	"errors"
	"gorm.io/gorm"
	"shop/pkg/model"
	"shop/pkg/repository"
	utils "shop/pkg/utils/token"
)

type authService struct {
	userRepository repository.UserRepository
}

func (a *authService) Login(au *model.LoginUser) (interface{}, error) {
	response := make(map[string]interface{})
	//检查用户是否存在
	data, err := a.userRepository.GetUserByAccount(au.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	//检查密码是否正确
	if data.PassWord != au.Password {
		return nil, errors.New("密码错误")
	}
	//生成token
	token, err := utils.GenerateToken(data.Account)
	response = map[string]interface{}{
		"token":   token,
		"message": "登录成功",
	}
	return response, nil
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}
