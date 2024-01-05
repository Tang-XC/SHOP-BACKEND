package service

import (
	"shop/pkg/model"
	"shop/pkg/repository"
	"time"
)

type productService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
	userRepository     repository.UserRepository
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
	//查找产品
	//product, err := p.productRepository.GetProductByID(id)
	//if err != nil {
	//	return err
	//}
	//查找用户
	//user, err := p.userRepository.GetUserByID(product.Owner)
	//if err != nil {
	//	return err
	//}
	//删除相关文件
	//bucketName := user.Name + user.Account + strconv.Itoa(int(user.ID))
	//common.DeleteFile(bucketName,)
	return p.productRepository.Delete(product)
}
func (p *productService) GetProductByID(id uint) (*model.Product, error) {
	return p.productRepository.GetProductByID(id)
}
func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository, userRepository repository.UserRepository) ProductService {
	return &productService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		userRepository:     userRepository,
	}
}
