package service

import (
	"fmt"
	"shop/pkg/model"
	"shop/pkg/repository"
	"time"
)

type productService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
	userRepository     repository.UserRepository
	fileRepository     repository.FileRepository
	fileService        FileService
}

func (p *productService) List(page int, size int, category int) (model.ProductsResponse, error) {
	return p.productRepository.List(page, size, category)
}
func (p *productService) Create(addProduct *model.AddProduct, user *model.User) (string, error) {
	message := "创建成功"
	//查询分类是否存在
	if _, err := p.categoryRepository.GetCategoryByID(addProduct.Category); err != nil {
		return "", err
	}
	addProduct.CreatedAt = time.Now()
	addProduct.Owner = user.ID
	product := addProduct.GetProduct()

	//创建商品
	productResponse, err := p.productRepository.Create(product)
	if err != nil {
		return "", err
	}

	//添加文件
	fmt.Println()
	for _, file := range addProduct.Files {
		if err := p.AddFile(productResponse.ID, file); err != nil {
			return "", err
		}
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
	product, err := p.productRepository.GetProductByID(id)
	files := product.Files
	for _, file := range product.Files {
		files = append(files, file)
	}
	if err != nil {
		return err
	}

	if err := p.productRepository.Delete(product); err != nil {
		return err
	}
	for _, file := range files {
		if err := p.fileService.Delete(&file); err != nil {
			return err
		}
	}
	return nil
}
func (p *productService) GetProductByID(id uint) (*model.Product, error) {
	return p.productRepository.GetProductByID(id)
}
func (p *productService) AddFile(productId uint, fileId uint) error {
	product, err := p.GetProductByID(productId)
	if err != nil {
		return err
	}
	file, err := p.fileRepository.GetFileByID(fileId)
	if err != nil {
		return err
	}
	return p.productRepository.AddFile(product, file)
}
func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository, userRepository repository.UserRepository, fileRepository repository.FileRepository, fileService FileService) ProductService {
	return &productService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		userRepository:     userRepository,
		fileRepository:     fileRepository,
		fileService:        fileService,
	}
}
