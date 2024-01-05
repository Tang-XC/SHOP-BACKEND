package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/service"
)

type UploadController struct {
	uploadService service.UploadService
}

func (f *UploadController) UploadImages(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	token := c.GetHeader("Authorization")
	result, err := f.uploadService.UploadImages(files, token)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, result)
}

func (f *UploadController) RegisterRoute(api *gin.RouterGroup) {
	v1 := api.Group("/", middleware.Auth())
	{
		v1.POST("/uploadImages", f.UploadImages)
	}
}

func (f *UploadController) Name() string {
	return "File"
}
func NewUploadController(uploadService service.UploadService) *UploadController {
	return &UploadController{uploadService: uploadService}
}
