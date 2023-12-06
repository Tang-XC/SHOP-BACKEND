package controller

import (
	"github.com/gin-gonic/gin"
	"shop/pkg/common"
	"shop/pkg/model"
	"shop/pkg/service"
)

type AuthController struct {
	userService service.UserService
	authService service.AuthService
}

func (a *AuthController) Login(c *gin.Context) {
	loginInfo := new(model.LoginUser)
	if err := c.BindJSON(loginInfo); err != nil {
		common.FailedResponse(c, 401, err)
		return
	}
	result, err := a.authService.Login(loginInfo)
	if err != nil {
		common.FailedResponse(c, 401, err)
		return
	}
	common.SuccessResponse(c, result)
}
func (a *AuthController) Register(c *gin.Context) {
	registerInfo := new(model.RegisterUser)
	if err := c.BindJSON(registerInfo); err != nil {
		common.FailedResponse(c, 400, err)
		return
	}
	user := registerInfo.GetUser()
	result, err := a.userService.Create(user)
	if err != nil {
		common.FailedResponse(c, 400, err)
		return
	}
	common.SuccessResponse(c, result)
}
func (a *AuthController) Name() string {
	return "auth"
}
func (a *AuthController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/login", a.Login)
	api.POST("/register", a.Register)
}
func NewAuthController(u service.UserService, a service.AuthService) Controller {
	return &AuthController{
		userService: u,
		authService: a,
	}
}
