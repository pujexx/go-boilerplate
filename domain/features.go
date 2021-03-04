package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Features struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name      string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"required,max=255"`
	Path      string    `gorm:"column:path;size:255;type:varchar;" json:"path" validate:"required,max=255"`
	Method    string    `gorm:"column:method;size:255;type:varchar;" json:"method" validate:"required,max=255"`
}

func (Features) TableName() string {
	return "features"
}

// Interface Repository
type FeaturesRepository interface {
	Save(p *Features) error
	Update(p *Features) error
	ByID(id string) (Features, error)
	Find(page int, filters []lib.FilterDomain) ([]Features, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Features
	Delete(id string) error
}

// Interface Service

type FeaturesService interface {
	Save(p *Features) error
	Update(p *Features) error
	ByID(id string) (Features, error)
	Find(page int, filters []lib.FilterDomain) ([]Features, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Features)
	Delete(id string) error
}
