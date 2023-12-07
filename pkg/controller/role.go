package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/model"
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

func (r *RoleController) Create(c *gin.Context) {
	roleAdd := new(model.AddRole)
	if err := c.BindJSON(roleAdd); err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	role := roleAdd.GetRole()
	result, err := r.roleService.Create(role)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, result)
}

func (r *RoleController) Name() string {
	return "User"
}

func (r *RoleController) RegisterRoute(api *gin.RouterGroup) {
	v1 := api.Group("/", middleware.Auth())
	{
		v1.GET("/roles", r.List)
		v1.POST("/role", r.Create)
	}
}
func NewRoleController(roleService service.RoleService) Controller {
	return &RoleController{roleService: roleService}
}
