package service

import (
	"shop/pkg/model"
	"shop/pkg/repository"
)

type productService struct {
	productRepository repository.ProductRepository
}

func (p *productService) List() (model.Products, error) {
	return p.productRepository.List()
}
func (p *productService) Create(product *model.Product) (string, error) {
	message := "创建成功"
	_, err := p.productRepository.Create(product)
	if err != nil {
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
func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}
