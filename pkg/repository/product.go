package repository

import (
	"gorm.io/gorm"
	"shop/pkg/model"
)

type productRepository struct {
	db *gorm.DB
}

func (u *productRepository) List(page int, size int, category int, keywords string) (model.ProductsResponse, error) {
	offset := page * size
	products := make(model.Products, 0)
	var total int64 = 0
	result := model.ProductsResponse{
		Total:    total,
		Products: products,
	}
	//根据关键字查询

	//根据分类查询
	if category != 0 {
		if err := u.db.Order("created_at desc").Where("category = ?", category).Offset(offset).Limit(size).Preload("Files").Find(&products).Error; err != nil {
			return result, err
		}
		if err := u.db.Model(&model.Product{}).Where("category = ?", category).Count(&total).Error; err != nil {
			return result, err
		}
	} else {
		if err := u.db.Order("created_at desc").Offset(offset).Limit(size).Preload("Files").Find(&products).Error; err != nil {
			return result, err
		}
		if err := u.db.Model(&model.Product{}).Count(&total).Error; err != nil {
			return result, err
		}
	}

	result.Total = total
	result.Products = products
	return result, nil
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
	if err := u.db.Model(&product).Association("Files").Clear(); err != nil {
		return err
	}
	if err := u.db.Delete(product).Error; err != nil {
		return err
	}
	return nil
}
func (u *productRepository) RemoveFileRelation(product *model.Product) error {
	if err := u.db.Model(&product).Association("Files").Clear(); err != nil {
		return err
	}
	return nil
}
func (u *productRepository) GetProductByID(id uint) (*model.Product, error) {
	//product := new(model.Product)
	//if err := u.db.First(product, id).Error; err != nil {
	//	return nil, err
	//}
	//return product, nil
	product := new(model.Product)
	if err := u.db.Preload("Files").First(product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (u *productRepository) AddFile(product *model.Product, file *model.File) error {
	if err := u.db.Model(product).Association("Files").Append(file); err != nil {
		return err
	}
	return nil
}
func (u *productRepository) Migrate() error {
	return u.db.AutoMigrate(&model.Product{})
}
func newProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
