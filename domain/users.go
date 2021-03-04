package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Users struct {
	Id        string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Username  string    `gorm:"column:username;size:255;type:varchar;" json:"username" validate:"required,max=255"`
	Password  string    `gorm:"column:password;size:255;type:varchar;" json:"password" validate:"required,max=255"`
	Email     string    `gorm:"column:email;size:255;type:varchar;" json:"email" validate:"required,max=255,email"`
	FullName  string    `gorm:"column:full_name;size:255;type:varchar;" json:"full_name" validate:"required,max=255"`
	Avatar    string    `gorm:"column:avatar;size:255;type:varchar;" json:"avatar" validate:"max=255"`
}

func (Users) TableName() string {
	return "users"
}

// Interface Repository
type UsersRepository interface {
	Save(p *Users) error
	Update(p *Users) error
	ByID(id string) (Users, error)
	Find(page int, filters []lib.FilterDomain) ([]Users, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Users
	Delete(id string) error
}

// Interface Service

type UsersService interface {
	Save(p *Users) error
	Update(p *Users) error
	ByID(id string) (Users, error)
	Find(page int, filters []lib.FilterDomain) ([]Users, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Users)
	Delete(id string) error
}
