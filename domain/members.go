package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Members struct {
	Id         string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Identity   string    `gorm:"column:identity;size:255;type:varchar;" json:"identity" validate:"max=255"`
	Name       string    `gorm:"column:name;size:255;type:varchar;" json:"name" validate:"max=255"`
	Phone      string    `gorm:"column:phone;size:255;type:varchar;" json:"phone" validate:"max=255"`
	TypeMember string    `gorm:"column:type_member;size:255;type:varchar;" json:"type_member" validate:"max=255"`
	Address    string    `gorm:"column:address;size:255;type:varchar;" json:"address" validate:"max=255"`
}

func (Members) TableName() string {
	return "members"
}

// Interface Repository
type MembersRepository interface {
	Save(p *Members) error
	Update(p *Members) error
	ByID(id string) (Members, error)
	Find(page int, filters []lib.FilterDomain) ([]Members, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Members
	Delete(id string) error
}

// Interface Service

type MembersService interface {
	Save(p *Members) error
	Update(p *Members) error
	ByID(id string) (Members, error)
	Find(page int, filters []lib.FilterDomain) ([]Members, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Members)
	Delete(id string) error
}
