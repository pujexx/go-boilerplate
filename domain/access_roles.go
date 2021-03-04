package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type AccessRoles struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	RoleId    string    `gorm:"column:role_id;size:255;type:varchar;" json:"role_id" validate:"max=255"`
	FeatureId string    `gorm:"column:feature_id;size:255;type:varchar;" json:"feature_id" validate:"max=255"`
}

func (AccessRoles) TableName() string {
	return "access_roles"
}

// Interface Repository
type AccessRolesRepository interface {
	Save(p *AccessRoles) error
	Update(p *AccessRoles) error
	ByID(id string) (AccessRoles, error)
	Find(page int, filters []lib.FilterDomain) ([]AccessRoles, lib.Paginator)
	FindRange(from time.Time, to time.Time) []AccessRoles
	Delete(id string) error
}

// Interface Service

type AccessRolesService interface {
	Save(p *AccessRoles) error
	Update(p *AccessRoles) error
	ByID(id string) (AccessRoles, error)
	Find(page int, filters []lib.FilterDomain) ([]AccessRoles, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]AccessRoles)
	Delete(id string) error
}
