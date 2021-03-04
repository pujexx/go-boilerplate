package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Variants struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name      string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"required,max=255"`
	ProductId string    `gorm:"column:product_id;size:255;type:varchar;" json:"product_id" validate:"max=255"`
	Position  int       `gorm:"column:position;size:11;type:int;" json:"position" validate:"required,number"`
}

func (Variants) TableName() string {
	return "variants"
}

// Interface Repository
type VariantsRepository interface {
	Save(p *Variants) error
	Update(p *Variants) error
	ByID(id string) (Variants, error)
	Find(page int, filters []lib.FilterDomain) ([]Variants, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Variants
	Delete(id string) error
}

// Interface Service

type VariantsService interface {
	Save(p *Variants) error
	Update(p *Variants) error
	ByID(id string) (Variants, error)
	Find(page int, filters []lib.FilterDomain) ([]Variants, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Variants)
	Delete(id string) error
}
