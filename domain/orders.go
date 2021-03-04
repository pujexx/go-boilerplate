package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Orders struct {
	Id             string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	TransactionId  string    `gorm:"column:transaction_id;size:255;type:varchar;" json:"transaction_id" validate:"max=255"`
	ProductId      string    `gorm:"column:product_id;size:255;type:varchar;" json:"product_id" validate:"max=255"`
	Qty            int       `gorm:"column:qty;size:11;type:int;" json:"qty" validate:"number"`
	Note           string    `gorm:"column:note;size:255;type:varchar;" json:"note" validate:"max=255"`
	ProductName    string    `gorm:"column:product_name;size:255;type:varchar;" json:"product_name" validate:"max=255"`
	ProductExtra   string    `gorm:"column:product_extra;size:255;type:varchar;" json:"product_extra" validate:"max=255"`
	ProductVariant string    `gorm:"column:product_variant;size:255;type:varchar;" json:"product_variant" validate:"max=255"`
}

func (Orders) TableName() string {
	return "orders"
}

// Interface Repository
type OrdersRepository interface {
	Save(p *Orders) error
	Update(p *Orders) error
	ByID(id string) (Orders, error)
	Find(page int, filters []lib.FilterDomain) ([]Orders, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Orders
	Delete(id string) error
}

// Interface Service

type OrdersService interface {
	Save(p *Orders) error
	Update(p *Orders) error
	ByID(id string) (Orders, error)
	Find(page int, filters []lib.FilterDomain) ([]Orders, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Orders)
	Delete(id string) error
}
