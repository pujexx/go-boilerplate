package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Stores struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name      string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"required,max=255"`
	Image     string    `gorm:"column:image;size:255;type:varchar;" json:"image" validate:"required,max=255"`
	Address   string    `gorm:"column:address;size:255;type:varchar;" json:"address" validate:"max=255"`
	UserId    string    `gorm:"column:user_id;size:255;type:varchar;" json:"user_id" validate:"max=255"`
}

func (Stores) TableName() string {
	return "stores"
}

// Interface Repository
type StoresRepository interface {
	Save(p *Stores) error
	Update(p *Stores) error
	ByID(id string) (Stores, error)
	Find(page int, filters []lib.FilterDomain) ([]Stores, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Stores
	Delete(id string) error
}

// Interface Service

type StoresService interface {
	Save(p *Stores) error
	Update(p *Stores) error
	ByID(id string) (Stores, error)
	Find(page int, filters []lib.FilterDomain) ([]Stores, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Stores)
	Delete(id string) error
}
