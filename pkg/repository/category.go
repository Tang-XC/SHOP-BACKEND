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

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}
