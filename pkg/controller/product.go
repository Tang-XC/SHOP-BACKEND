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
	userService    service.UserService
}

func (u *ProductController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	category, _ := strconv.Atoi(c.DefaultQuery("category", "0"))
	products, err := u.productService.List(page, size, category)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, products)
}
func (u *ProductController) Create(c *gin.Context) {
	//获取参数
	var addProduct = new(model.AddProduct)
	if err := c.ShouldBindJSON(&addProduct); err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	//获取用户信息
	token := c.GetHeader("Authorization")
	user, err := u.userService.GetUser(token)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	message, err := u.productService.Create(addProduct, user)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	product, err := u.productService.GetProductByID(uint(id))
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, product)
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
func NewProductController(productService service.ProductService, userService service.UserService) *ProductController {
	return &ProductController{productService: productService, userService: userService}
}
