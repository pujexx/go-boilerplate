package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type VariantValues struct {
	VariantOptionsId1 string    `gorm:"column:variant_options_id1;size:255;type:varchar;" json:"variant_options_id1" validate:"required,max=255"`
	VariantOptionsId2 string    `gorm:"column:variant_options_id2;size:255;type:varchar;" json:"variant_options_id2" validate:"max=255"`
	Stock             int       `gorm:"column:stock;size:11;type:int;" json:"stock" validate:"number"`
	Price             int       `gorm:"column:price;size:11;type:int;" json:"price" validate:"number"`
	Cogs              int       `gorm:"column:cogs;size:11;type:int;" json:"cogs" validate:"number"`
	Sku               string    `gorm:"column:sku;size:255;type:varchar;" json:"sku" validate:"max=255"`
	ProductId         string    `gorm:"column:product_id;size:255;type:varchar;" json:"product_id" validate:"max=255"`
	Id                string    `gorm:"column:id;size:200;type:varchar;" json:"id" validate:"max=200"`
	CreatedAt         time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt         time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
}

func (VariantValues) TableName() string {
	return "variant_values"
}

// Interface Repository
type VariantValuesRepository interface {
	Save(p *VariantValues) error
	Update(p *VariantValues) error
	ByID(id string) (VariantValues, error)
	Find(page int, filters []lib.FilterDomain) ([]VariantValues, lib.Paginator)
	FindRange(from time.Time, to time.Time) []VariantValues
	Delete(id string) error
}

// Interface Service

type VariantValuesService interface {
	Save(p *VariantValues) error
	Update(p *VariantValues) error
	ByID(id string) (VariantValues, error)
	Find(page int, filters []lib.FilterDomain) ([]VariantValues, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]VariantValues)
	Delete(id string) error
}
