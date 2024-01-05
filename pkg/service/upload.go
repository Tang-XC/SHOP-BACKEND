package service

import (
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"shop/pkg/common"
	"shop/pkg/config"
	"shop/pkg/model"
	"shop/pkg/repository"
	utils "shop/pkg/utils/token"
	"strconv"
	"time"
)

type uploadService struct {
	minioClient    *minio.Client
	config         config.MinioConfig
	userRepository repository.UserRepository
}

func (f *uploadService) UploadImages(files []*multipart.FileHeader, token string) ([]model.FileResponse, error) {
	//声明用于存放多个FileResponse类型的变量
	var fr []model.FileResponse = make([]model.FileResponse, 0)

	//根据token获取用户信息
	claim, err := utils.ParseToken(token)
	if err != nil {
		return fr, err
	}
	user, err := f.userRepository.GetUserByAccount(claim.Subject)
	if err != nil {
		return fr, err
	}
	//创建用户的bucket
	bucketName := user.Name + user.Account + strconv.Itoa(int(user.ID))
	common.CreateBucket(bucketName, f.minioClient)
	common.CreateFolder(bucketName, "images", f.minioClient)

	//上传文件
	for _, file := range files {
		//为文件名称加上时间戳
		time := strconv.Itoa(int(time.Now().Unix()))
		fileName := "images/" + time + file.Filename
		src, err := file.Open()
		if err != nil {
			return fr, err
		}
		defer src.Close()

		//开始上传
		_, err = common.FileUploader(bucketName, fileName, src, file.Size, f.minioClient)
		if err != nil {
			return fr, err
		}
		//获取文件的信息
		fileInfo, err := common.FileInfo(bucketName, fileName, f.minioClient)

		//获取文件的url
		result, err := common.FileURL(bucketName, fileName, f.minioClient)
		if err != nil {
			return fr, err
		}
		fr = append(fr, model.FileResponse{
			Uid:    fileInfo.ETag,
			Name:   file.Filename,
			Status: "done",
			Url:    result,
			Size:   fileInfo.Size,
		})
	}
	return fr, nil
}

func NewUploadService(minioClient *minio.Client, config config.MinioConfig, userRepository repository.UserRepository) UploadService {
	return &uploadService{
		minioClient:    minioClient,
		config:         config,
		userRepository: userRepository,
	}
}
