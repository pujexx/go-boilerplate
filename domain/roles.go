package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Roles struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Name      string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"required,max=255"`
}

func (Roles) TableName() string {
	return "roles"
}

// Interface Repository
type RolesRepository interface {
	Save(p *Roles) error
	Update(p *Roles) error
	ByID(id string) (Roles, error)
	Find(page int, filters []lib.FilterDomain) ([]Roles, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Roles
	Delete(id string) error
}

// Interface Service

type RolesService interface {
	Save(p *Roles) error
	Update(p *Roles) error
	ByID(id string) (Roles, error)
	Find(page int, filters []lib.FilterDomain) ([]Roles, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Roles)
	Delete(id string) error
}
