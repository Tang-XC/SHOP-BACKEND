package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/model"
	"shop/pkg/service"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func (cg *CategoryController) List(c *gin.Context) {
	categorys, err := cg.categoryService.List()
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, categorys)
}
func (cg *CategoryController) Create(c *gin.Context) {
	var addCategory = new(model.AddCategory)
	if err := c.ShouldBindJSON(&addCategory); err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	category := addCategory.GetCategory()
	message, err := cg.categoryService.Create(&category)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, message)
}

func (u *CategoryController) RegisterRoute(api *gin.RouterGroup) {
	v1 := api.Group("/", middleware.Auth())
	{
		v1.GET("/categories", u.List)
		v1.POST("/category", u.Create)
	}
}

func (u *CategoryController) Name() string {
	return "Category"
}
func NewCategoryController(categoryService service.CategoryService) Controller {
	return &CategoryController{categoryService}
}
