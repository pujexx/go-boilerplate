package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Cdns struct {
	Id            string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	PathOri       string    `gorm:"column:path_ori;size:255;type:varchar;" json:"path_ori" validate:"max=255"`
	PathThumbnail string    `gorm:"column:path_thumbnail;size:255;type:varchar;" json:"path_thumbnail" validate:"max=255"`
	StoreId       string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
}

func (Cdns) TableName() string {
	return "cdns"
}

// Interface Repository
type CdnsRepository interface {
	Save(p *Cdns) error
	Update(p *Cdns) error
	ByID(id string) (Cdns, error)
	Find(page int, filters []lib.FilterDomain) ([]Cdns, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Cdns
	Delete(id string) error
}

// Interface Service

type CdnsService interface {
	Save(p *Cdns) error
	Update(p *Cdns) error
	ByID(id string) (Cdns, error)
	Find(page int, filters []lib.FilterDomain) ([]Cdns, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Cdns)
	Delete(id string) error
}
