package controller

import (
	"github.com/gin-gonic/gin"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/service"
)

type RoleController struct {
	roleService service.RoleService
}

func (r *RoleController) List(c *gin.Context) {
	roles, err := r.roleService.List()
	if err != nil {
		common.FailedResponse(c, 400, err)
	}
	common.SuccessResponse(c, roles)
}

func (r *RoleController) Name() string {
	return "User"
}

func (r *RoleController) RegisterRoute(api *gin.RouterGroup) {
	v1 := api.Group("/", middleware.Auth())
	{
		v1.GET("/roles", r.List)
	}

}
func NewRoleController(roleService service.RoleService) Controller {
	return &RoleController{roleService: roleService}
}
