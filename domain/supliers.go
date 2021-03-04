package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Supliers struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name      string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"max=255"`
	Address   string    `gorm:"column:address;size:255;type:varchar;" json:"address" validate:"max=255"`
	Phone     string    `gorm:"column:phone;size:255;type:varchar;" json:"phone" validate:"max=255"`
	Pic       string    `gorm:"column:pic;size:255;type:varchar;" json:"pic" validate:"max=255"`
	StoreId   string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
}

func (Supliers) TableName() string {
	return "supliers"
}

// Interface Repository
type SupliersRepository interface {
	Save(p *Supliers) error
	Update(p *Supliers) error
	ByID(id string) (Supliers, error)
	Find(page int, filters []lib.FilterDomain) ([]Supliers, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Supliers
	Delete(id string) error
}

// Interface Service

type SupliersService interface {
	Save(p *Supliers) error
	Update(p *Supliers) error
	ByID(id string) (Supliers, error)
	Find(page int, filters []lib.FilterDomain) ([]Supliers, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Supliers)
	Delete(id string) error
}
