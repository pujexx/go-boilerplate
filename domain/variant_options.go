package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type VariantOptions struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name      string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"required,max=255"`
	VariantId string    `gorm:"column:variant_id;size:255;type:varchar;" json:"variant_id" validate:"max=255"`
	Slug      string    `gorm:"column:slug;size:255;type:varchar;" json:"slug" validate:"required,max=255"`
}

func (VariantOptions) TableName() string {
	return "variant_options"
}

// Interface Repository
type VariantOptionsRepository interface {
	Save(p *VariantOptions) error
	Update(p *VariantOptions) error
	ByID(id string) (VariantOptions, error)
	Find(page int, filters []lib.FilterDomain) ([]VariantOptions, lib.Paginator)
	FindRange(from time.Time, to time.Time) []VariantOptions
	Delete(id string) error
}

// Interface Service

type VariantOptionsService interface {
	Save(p *VariantOptions) error
	Update(p *VariantOptions) error
	ByID(id string) (VariantOptions, error)
	Find(page int, filters []lib.FilterDomain) ([]VariantOptions, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]VariantOptions)
	Delete(id string) error
}
