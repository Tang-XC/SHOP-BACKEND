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

type fileService struct {
	fileRepository repository.FileRepository
	minioClient    *minio.Client
	config         config.MinioConfig
	userRepository repository.UserRepository
}

func (f fileService) List() (model.Files, error) {
	return f.fileRepository.List()
}
func (f fileService) Create(files []*multipart.FileHeader, token string) (model.Files, error) {
	var filesData model.Files

	//根据token获取用户信息
	claim, err := utils.ParseToken(token)
	if err != nil {
		return filesData, err
	}
	user, err := f.userRepository.GetUserByAccount(claim.Subject)
	if err != nil {
		return filesData, err
	}
	//创建用户的bucket
	bucketName := user.Name + user.Account + strconv.Itoa(int(user.ID))
	common.CreateBucket(bucketName, f.minioClient)
	common.CreateFolder(bucketName, "images", f.minioClient)
	//上传文件
	for _, file := range files {
		//为文件名称加上时间戳
		timeStamp := strconv.Itoa(int(time.Now().Unix()))
		fileName := timeStamp + file.Filename
		path := "images/" + fileName

		src, err := file.Open()
		if err != nil {
			return filesData, err
		}
		defer src.Close()

		//开始上传
		_, err = common.FileUploader(bucketName, path, src, file.Size, f.minioClient)
		if err != nil {
			return filesData, err
		}
		//获取文件的信息
		fileInfo, err := common.FileInfo(bucketName, path, f.minioClient)

		//获取文件的url
		result := common.FileURL(bucketName, path, f.minioClient)
		if err != nil {
			return filesData, err
		}

		//将文件信息存入数据库
		addFile := model.AddFile{
			Uid:        fileInfo.ETag,
			Name:       fileName,
			Url:        result,
			Size:       file.Size,
			FileType:   file.Header.Get("Content-Type"),
			CreatedAt:  time.Now().Unix(),
			Path:       path,
			BucketName: bucketName,
		}
		data := addFile.GetFile()
		theFile, err := f.fileRepository.Create(data)
		filesData = append(filesData, *theFile)
		if err != nil {
			return filesData, err
		}
	}
	return filesData, err
}

func (f fileService) Delete(file *model.File) error {
	common.DeleteFile(file.BucketName, file.Path, f.minioClient)
	return f.fileRepository.Delete(file)
}

func (f fileService) GetFileByID(u uint) (*model.File, error) {
	return f.fileRepository.GetFileByID(u)
}

func NewFileService(fileRepository repository.FileRepository, minioClient *minio.Client, config config.MinioConfig, userRepository repository.UserRepository) FileService {
	return &fileService{
		fileRepository: fileRepository,
		minioClient:    minioClient,
		config:         config,
		userRepository: userRepository,
	}
}
