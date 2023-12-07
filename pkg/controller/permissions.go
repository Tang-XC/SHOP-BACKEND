package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/model"
	"shop/pkg/service"
)

type PermissionController struct {
	permissionService service.PermissionService
}

func (p *PermissionController) List(c *gin.Context) {
	permissions, err := p.permissionService.List()
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, permissions)
}
func (p *PermissionController) Create(c *gin.Context) {
	permissionAdd := new(model.AddPermission)
	if err := c.BindJSON(permissionAdd); err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	permission := permissionAdd.GetPermission()
	result, err := p.permissionService.Create(permission)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, result)
}

func (p *PermissionController) Name() string {
	return "permission"
}
func (p *PermissionController) RegisterRoute(api *gin.RouterGroup) {
	v1 := api.Group("/", middleware.Auth())
	{
		v1.GET("/permissions", p.List)
		v1.POST("/permission", p.Create)
	}
}

func NewPermissionController(p service.PermissionService) Controller {
	return &PermissionController{
		permissionService: p,
	}
}
