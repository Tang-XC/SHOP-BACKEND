package model

import "time"

type Product struct {
	ID        uint           `gorm:";primary_key;column:id" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Price     int            `gorm:"column:price" json:"price"`
	Brand     string         `gorm:"column:brand" json:"brand"`
	Image     string         `gorm:"column:image" json:"image"`
	Desc      string         `gorm:"column:desc" json:"desc"`
	Stock     int            `gorm:"column:stock" json:"stock"`
	Category  uint           `gorm:"column:category" json:"category"`
	Specs     Specifications `gorm:"embedded" json:"specs"`
	CreatedAt int64          `gorm:"column:created_at" json:"created_at"`
	Owner     uint           `gorm:"column:owner" json:"owner"`
	Files     Files          `gorm:"many2many:product_files;" json:"files"`
}

func (p Product) TableName() string {
	return "products"
}

type Products []Product
type ProductsResponse struct {
	Total    int64    `json:"total"`
	Products Products `json:"products"`
	Category Category `json:"category"`
}

type Specifications struct {
	ScreenSize     string
	Resolution     string
	Processor      string
	Storage        string
	RAM            string
	FrontCamera    string
	RearCamera     string
	Battery        string
	OS             string
	NetworkSupport string
}

type AddProduct struct {
	Name      string         `json:"name" `
	Price     int            `json:"price" `
	Brand     string         `json:"brand"`
	Image     string         `json:"image" `
	Desc      string         `json:"desc"`
	Stock     int            `json:"stock"`
	Category  uint           `json:"category"`
	Specs     Specifications `json:"specs"`
	CreatedAt time.Time      `json:"created_at"`
	Owner     uint           `json:"owner"`
	Files     []uint         `json:"files"`
}

func (a AddProduct) GetProduct() *Product {
	return &Product{
		Name:      a.Name,
		Price:     a.Price,
		Brand:     a.Brand,
		Image:     a.Image,
		Desc:      a.Desc,
		Stock:     a.Stock,
		Category:  a.Category,
		Specs:     a.Specs,
		CreatedAt: a.CreatedAt.Unix(),
		Owner:     a.Owner,
	}
}
