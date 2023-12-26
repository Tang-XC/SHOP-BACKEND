package service

import (
	"fmt"
	"shop/pkg/model"
	"shop/pkg/repository"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func (c *categoryService) List() (model.Categorys, error) {
	return c.categoryRepository.List()
}

// Create(*model.Category) (string, error)
func (c *categoryService) Create(category *model.Category) (string, error) {
	fmt.Println(category)
	category, err := c.categoryRepository.Create(category)
	if err != nil {
		return "", err
	}
	return "添加成功", nil
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}
