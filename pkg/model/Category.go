package model

type Category struct {
	ID   uint   `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (c Category) TableName() string {
	return "categories"
}

type Categorys []Category

type AddCategory struct {
	Name string `json:"name"`
}

func (a AddCategory) GetCategory() Category {
	return Category{
		Name: a.Name,
	}
}
