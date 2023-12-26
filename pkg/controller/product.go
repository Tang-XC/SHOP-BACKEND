package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/model"
	"shop/pkg/service"
	"strconv"
)

type ProductController struct {
	productService service.ProductService
}

func (u *ProductController) List(c *gin.Context) {
	products, err := u.productService.List()
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, products)
}
func (u *ProductController) Create(c *gin.Context) {
	var addProduct = new(model.AddProduct)
	if err := c.ShouldBindJSON(&addProduct); err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	message, err := u.productService.Create(addProduct)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, message)
}
func (u *ProductController) Update(c *gin.Context) {

}
func (u *ProductController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	if err := u.productService.Delete(uint(id)); err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, "删除成功")
}
func (u *ProductController) GetProductByID(c *gin.Context) {

}
func (u *ProductController) RegisterRoute(api *gin.RouterGroup) {
	v1 := api.Group("/", middleware.Auth())
	{
		v1.GET("/products", u.List)
		v1.GET("/product/:id", u.GetProductByID)
		v1.POST("/product", u.Create)
		v1.PUT("/product/:id", u.Update)
		v1.DELETE("/product/:id", u.Delete)
	}
}
func (u *ProductController) Name() string {
	return "Product"
}
func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}
