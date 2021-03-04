package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Categories struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name      string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"required,max=255"`
	StoreId   string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
	IsExtras  int       `gorm:"column:is_extras;size:1;type:tinyint;" json:"is_extras" validate:"number"`
}

func (Categories) TableName() string {
	return "categories"
}

// Interface Repository
type CategoriesRepository interface {
	Save(p *Categories) error
	Update(p *Categories) error
	ByID(id string) (Categories, error)
	Find(page int, filters []lib.FilterDomain) ([]Categories, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Categories
	Delete(id string) error
}

// Interface Service

type CategoriesService interface {
	Save(p *Categories) error
	Update(p *Categories) error
	ByID(id string) (Categories, error)
	Find(page int, filters []lib.FilterDomain) ([]Categories, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Categories)
	Delete(id string) error
}
