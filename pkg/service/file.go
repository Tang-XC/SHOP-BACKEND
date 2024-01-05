package service

import (
	"shop/pkg/model"
	"shop/pkg/repository"
)

type fileService struct {
	fileRepository repository.FileRepository
}

func (f fileService) List() (model.Files, error) {
	return f.fileRepository.List()
}
func (f fileService) Create(file *model.File) (*model.File, error) {
	return f.fileRepository.Create(file)
}

func (f fileService) Delete(file *model.File) error {
	return f.fileRepository.Delete(file)
}

func (f fileService) GetFileByID(u uint) (*model.File, error) {
	return f.fileRepository.GetFileByID(u)
}

func NewFileService(fileRepository repository.FileRepository) FileService {
	return &fileService{fileRepository}
}
