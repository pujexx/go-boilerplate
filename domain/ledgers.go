package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Ledgers struct {
	Id              string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	Description     string    `gorm:"column:description;size:255;type:varchar;" json:"description" validate:"max=255"`
	DateTransaction time.Time `gorm:"column:date_transaction;type:datetime;" json:"date_transaction"`
	Debit           int       `gorm:"column:debit;size:11;type:int;" json:"debit" validate:"number"`
	Credit          int       `gorm:"column:credit;size:11;type:int;" json:"credit" validate:"number"`
	Amount          int       `gorm:"column:amount;size:11;type:int;" json:"amount" validate:"number"`
	Ref             string    `gorm:"column:ref;size:255;type:varchar;" json:"ref" validate:"max=255"`
	Type            string    `gorm:"column:type;size:255;type:varchar;" json:"type" validate:"max=255"`
	StoreId         string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
}

func (Ledgers) TableName() string {
	return "ledgers"
}

// Interface Repository
type LedgersRepository interface {
	Save(p *Ledgers) error
	Update(p *Ledgers) error
	ByID(id string) (Ledgers, error)
	Find(page int, filters []lib.FilterDomain) ([]Ledgers, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Ledgers
	Delete(id string) error
}

// Interface Service

type LedgersService interface {
	Save(p *Ledgers) error
	Update(p *Ledgers) error
	ByID(id string) (Ledgers, error)
	Find(page int, filters []lib.FilterDomain) ([]Ledgers, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Ledgers)
	Delete(id string) error
}
