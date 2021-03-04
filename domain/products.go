package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Products struct {
	Id          string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name        string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"required,max=255"`
	Description string    `gorm:"column:description;size:255;type:varchar;" json:"description" validate:"max=255"`
	Price       int       `gorm:"column:price;size:11;type:int;" json:"price" validate:"number"`
	Cogs        int       `gorm:"column:cogs;size:11;type:int;" json:"cogs" validate:"number"`
	Stock       int       `gorm:"column:stock;size:11;type:int;" json:"stock" validate:"number"`
	Image       string    `gorm:"column:image;size:255;type:varchar;" json:"image" validate:"max=255"`
	CategoryId  string    `gorm:"column:category_id;size:255;type:varchar;" json:"category_id" validate:"max=255"`
	StoreId     string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
	Sku         string    `gorm:"column:sku;size:255;type:varchar;" json:"sku" validate:"max=255"`
	Brand       string    `gorm:"column:brand;size:255;type:varchar;" json:"brand" validate:"max=255"`
	FlagColor   string    `gorm:"column:flag_color;size:255;type:varchar;" json:"flag_color" validate:"max=255"`
}

func (Products) TableName() string {
	return "products"
}

// Interface Repository
type ProductsRepository interface {
	Save(p *Products) error
	Update(p *Products) error
	ByID(id string) (Products, error)
	Find(page int, filters []lib.FilterDomain) ([]Products, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Products
	Delete(id string) error
}

// Interface Service

type ProductsService interface {
	Save(p *Products) error
	Update(p *Products) error
	ByID(id string) (Products, error)
	Find(page int, filters []lib.FilterDomain) ([]Products, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Products)
	Delete(id string) error
}
