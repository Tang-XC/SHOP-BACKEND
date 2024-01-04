package service

import (
	"shop/pkg/model"
	"shop/pkg/repository"
)

type productService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
}

func (p *productService) List() (model.Products, error) {
	return p.productRepository.List()
}
func (p *productService) Create(addProduct *model.AddProduct) (string, error) {
	message := "创建成功"
	//查询分类是否存在
	if _, err := p.categoryRepository.GetCategoryByID(addProduct.Category); err != nil {
		return "", err
	}
	product := addProduct.GetProduct()
	//创建商品
	if _, err := p.productRepository.Create(product); err != nil {
		return "", err
	}
	return message, nil
}
func (p *productService) Update(id uint, product *model.Product) (*model.Product, error) {
	old, err := p.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	product.ID = old.ID
	return p.productRepository.Update(product)
}
func (p *productService) Delete(id uint) error {
	product := &model.Product{
		ID: id,
	}
	return p.productRepository.Delete(product)
}
func (p *productService) GetProductByID(id uint) (*model.Product, error) {
	return p.productRepository.GetProductByID(id)
}
func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository) ProductService {
	return &productService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
	}
}
