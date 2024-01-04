package model

type Product struct {
	ID       uint           `gorm:";primary_key;column:id" json:"id"`
	Name     string         `gorm:"column:name" json:"name"`
	Price    int            `gorm:"column:price" json:"price"`
	Brand    string         `gorm:"column:brand" json:"brand"`
	Image    string         `gorm:"column:image" json:"image"`
	Desc     string         `gorm:"column:desc" json:"desc"`
	Stock    int            `gorm:"column:stock" json:"stock"`
	Category uint           `gorm:"column:category" json:"category"`
	Specs    Specifications `gorm:"embedded" json:"specs"`
}

func (p Product) TableName() string {
	return "products"
}

type Products []Product

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
	Name     string         `json:"name" `
	Price    int            `json:"price" `
	Brand    string         `json:"brand"`
	Image    string         `json:"image" `
	Desc     string         `json:"desc"`
	Stock    int            `json:"stock"`
	Category uint           `json:"category"`
	Specs    Specifications `json:"specs"`
}

func (a AddProduct) GetProduct() *Product {
	return &Product{
		Name:     a.Name,
		Price:    a.Price,
		Brand:    a.Brand,
		Image:    a.Image,
		Desc:     a.Desc,
		Stock:    a.Stock,
		Category: a.Category,
		Specs:    a.Specs,
	}
}
