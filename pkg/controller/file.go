package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/common"
	"shop/pkg/middleware"
	"shop/pkg/service"
	"strconv"
)

type FileController struct {
	fileService service.FileService
}

func (f FileController) Name() string {
	return "File"
}
func (f FileController) List(c *gin.Context) {
	files, err := f.fileService.List()
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
	}
	common.SuccessResponse(c, files)
}
func (f FileController) Create(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	token := c.GetHeader("Authorization")
	filesResponse, err := f.fileService.Create(files, token)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, filesResponse)
}
func (f FileController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	file, err := f.fileService.GetFileByID(uint(id))
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	err = f.fileService.Delete(file)
	if err != nil {
		common.FailedResponse(c, http.StatusBadRequest, err)
		return
	}
	common.SuccessResponse(c, "删除成功")
}

func (f FileController) RegisterRoute(group *gin.RouterGroup) {
	v1 := group.Group("/", middleware.Auth())
	{
		v1.GET("/files", f.List)
		v1.POST("/file", f.Create)
		v1.DELETE("/file/:id", f.Delete)
	}
}

func NewFileController(fileService service.FileService) Controller {
	return &FileController{fileService: fileService}
}
