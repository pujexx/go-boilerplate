package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type UserRoles struct {
	StoreId   string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
	UserId    string    `gorm:"column:user_id;size:255;type:varchar;" json:"user_id" validate:"max=255"`
	RoleId    string    `gorm:"column:role_id;size:255;type:varchar;" json:"role_id" validate:"max=255"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
}

func (UserRoles) TableName() string {
	return "user_roles"
}

// Interface Repository
type UserRolesRepository interface {
	Save(p *UserRoles) error
	Update(p *UserRoles) error
	ByID(id string) (UserRoles, error)
	Find(page int, filters []lib.FilterDomain) ([]UserRoles, lib.Paginator)
	FindRange(from time.Time, to time.Time) []UserRoles
	Delete(id string) error
}

// Interface Service

type UserRolesService interface {
	Save(p *UserRoles) error
	Update(p *UserRoles) error
	ByID(id string) (UserRoles, error)
	Find(page int, filters []lib.FilterDomain) ([]UserRoles, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]UserRoles)
	Delete(id string) error
}
