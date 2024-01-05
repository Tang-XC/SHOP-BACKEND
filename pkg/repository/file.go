package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type fileRepository struct {
	db *gorm.DB
}

func (f *fileRepository) List() (model.Files, error) {
	var files model.Files
	err := f.db.Find(&files).Error
	return files, err
}
func (f *fileRepository) Create(file *model.File) (*model.File, error) {
	err := f.db.Create(file).Error
	return file, err
}
func (f *fileRepository) Delete(file *model.File) error {
	err := f.db.Delete(file).Error
	return err
}
func (f *fileRepository) Migrate() error {
	return f.db.AutoMigrate(&model.File{})
}
func (f *fileRepository) GetFileByID(id uint) (*model.File, error) {
	var file model.File
	err := f.db.First(&file, id).Error
	return &file, err
}
func newFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db}
}
