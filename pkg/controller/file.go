package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/service"
)

type FileController struct {
	fileService service.FileService
}

func (f *FileController) UploadImages(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	token := c.GetHeader("Authorization")
	result, err := f.fileService.UploadImages(files, token)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, result)
}

func (f *FileController) RegisterRoute(api *gin.RouterGroup) {
	v1 := api.Group("/", middleware.Auth())
	{
		v1.POST("/uploadImages", f.UploadImages)
	}
}

func (f *FileController) Name() string {
	return "File"
}
func NewFileController(fileService service.FileService) *FileController {
	return &FileController{fileService: fileService}
}
