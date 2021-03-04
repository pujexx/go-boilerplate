package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type ExtraOrders struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	ProductId string    `gorm:"column:product_id;size:255;type:varchar;" json:"product_id" validate:"max=255"`
	OrderId   string    `gorm:"column:order_id;size:255;type:varchar;" json:"order_id" validate:"max=255"`
}

func (ExtraOrders) TableName() string {
	return "extra_orders"
}

// Interface Repository
type ExtraOrdersRepository interface {
	Save(p *ExtraOrders) error
	Update(p *ExtraOrders) error
	ByID(id string) (ExtraOrders, error)
	Find(page int, filters []lib.FilterDomain) ([]ExtraOrders, lib.Paginator)
	FindRange(from time.Time, to time.Time) []ExtraOrders
	Delete(id string) error
}

// Interface Service

type ExtraOrdersService interface {
	Save(p *ExtraOrders) error
	Update(p *ExtraOrders) error
	ByID(id string) (ExtraOrders, error)
	Find(page int, filters []lib.FilterDomain) ([]ExtraOrders, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]ExtraOrders)
	Delete(id string) error
}
