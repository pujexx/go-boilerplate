package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type transactionsRepository struct {
	DB *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) domain.TransactionsRepository {
	return &transactionsRepository{
		DB: db,
	}
}

func (tR transactionsRepository) Save(t *domain.Transactions) error {
	err := tR.DB.Save(t).Error
	return err
}

func (tR transactionsRepository) Update(t *domain.Transactions) error {
	err := tR.DB.Save(t).Error
	return err
}

func (tR transactionsRepository) ByID(id string) (domain.Transactions, error) {
	var t domain.Transactions
	err := tR.DB.Where("id = ?", id).First(&t).Error
	return t, err
}
func (tR transactionsRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Transactions, lib.Paginator) {
	var ts []domain.Transactions
	db := tR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&ts, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return ts, paginate
}

func (tR transactionsRepository) FindRange(from time.Time, to time.Time) []domain.Transactions {
	var ts []domain.Transactions
	db := tR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&ts)
	return ts
}

func (tR transactionsRepository) Delete(id string) error {
	return tR.DB.Where("id = ?", id).Delete(&domain.Transactions{}).Error
}
