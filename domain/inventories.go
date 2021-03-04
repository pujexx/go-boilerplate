package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Inventories struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	ProductId string    `gorm:"column:product_id;size:255;type:varchar;" json:"product_id" validate:"max=255"`
	Stock     int       `gorm:"column:stock;size:11;type:int;" json:"stock" validate:"number"`
	Unit      string    `gorm:"column:unit;size:255;type:varchar;" json:"unit" validate:"max=255"`
	SuplierId string    `gorm:"column:suplier_id;size:255;type:varchar;" json:"suplier_id" validate:"max=255"`
	StoreId   string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
}

func (Inventories) TableName() string {
	return "inventories"
}

// Interface Repository
type InventoriesRepository interface {
	Save(p *Inventories) error
	Update(p *Inventories) error
	ByID(id string) (Inventories, error)
	Find(page int, filters []lib.FilterDomain) ([]Inventories, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Inventories
	Delete(id string) error
}

// Interface Service

type InventoriesService interface {
	Save(p *Inventories) error
	Update(p *Inventories) error
	ByID(id string) (Inventories, error)
	Find(page int, filters []lib.FilterDomain) ([]Inventories, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Inventories)
	Delete(id string) error
}
