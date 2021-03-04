package domain

import (
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type Transactions struct {
	Id            string    `gorm:"column:id;size:200;primaryKey;type:varchar;" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
	MemberId      string    `gorm:"column:member_id;size:255;type:varchar;" json:"member_id" validate:"max=255"`
	PaymentMethod string    `gorm:"column:payment_method;size:255;type:varchar;" json:"payment_method" validate:"max=255"`
	Amount        int       `gorm:"column:amount;size:11;type:int;" json:"amount" validate:"number"`
	UseTax        int       `gorm:"column:use_tax;size:1;type:tinyint;" json:"use_tax" validate:"number"`
	StoreId       string    `gorm:"column:store_id;size:255;type:varchar;" json:"store_id" validate:"max=255"`
}

func (Transactions) TableName() string {
	return "transactions"
}

// Interface Repository
type TransactionsRepository interface {
	Save(p *Transactions) error
	Update(p *Transactions) error
	ByID(id string) (Transactions, error)
	Find(page int, filters []lib.FilterDomain) ([]Transactions, lib.Paginator)
	FindRange(from time.Time, to time.Time) []Transactions
	Delete(id string) error
}

// Interface Service

type TransactionsService interface {
	Save(p *Transactions) error
	Update(p *Transactions) error
	ByID(id string) (Transactions, error)
	Find(page int, filters []lib.FilterDomain) ([]Transactions, lib.Paginator)
	FindRange(from time.Time, to time.Time) (error, *[]Transactions)
	Delete(id string) error
}
