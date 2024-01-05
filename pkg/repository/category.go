package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type categoryRepository struct {
	db *gorm.DB
}

func (c *categoryRepository) List() (model.Categorys, error) {
	var categorys model.Categorys
	err := c.db.Find(&categorys).Error
	if err != nil {
		return nil, err
	}
	return categorys, nil
}
func (c *categoryRepository) Create(category *model.Category) (*model.Category, error) {
	err := c.db.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
func (c *categoryRepository) Delete(id uint) error {
	return c.db.Delete(&model.Category{}, id).Error
}
func (c *categoryRepository) GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	err := c.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
func (c *categoryRepository) Migrate() error {
	return c.db.AutoMigrate(&model.Category{})
}

func newCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}
