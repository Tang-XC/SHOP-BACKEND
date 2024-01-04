package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type productRepository struct {
	db *gorm.DB
}

func (u *productRepository) List() (model.Products, error) {
	products := make(model.Products, 0)
	if er := u.db.Order("name").Find(&products).Error; er != nil {
		return nil, er
	}
	return products, nil
}
func (u *productRepository) Create(product *model.Product) (*model.Product, error) {
	if err := u.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (u *productRepository) Update(product *model.Product) (*model.Product, error) {
	if err := u.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (u *productRepository) Delete(product *model.Product) error {
	return u.db.Delete(product).Error
}
func (u *productRepository) GetProductByID(id uint) (*model.Product, error) {
	product := new(model.Product)
	if err := u.db.First(product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (u *productRepository) Migrate() error {
	return u.db.AutoMigrate(&model.Product{})
}
func newProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
