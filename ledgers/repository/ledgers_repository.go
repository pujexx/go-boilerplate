package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ledgersRepository struct {
	DB *gorm.DB
}

func NewLedgersRepository(db *gorm.DB) domain.LedgersRepository {
	return &ledgersRepository{
		DB: db,
	}
}

func (lR ledgersRepository) Save(l *domain.Ledgers) error {
	err := lR.DB.Save(l).Error
	return err
}

func (lR ledgersRepository) Update(l *domain.Ledgers) error {
	err := lR.DB.Save(l).Error
	return err
}

func (lR ledgersRepository) ByID(id string) (domain.Ledgers, error) {
	var l domain.Ledgers
	err := lR.DB.Where("id = ?", id).First(&l).Error
	return l, err
}
func (lR ledgersRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Ledgers, lib.Paginator) {
	var ls []domain.Ledgers
	db := lR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&ls, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return ls, paginate
}

func (lR ledgersRepository) FindRange(from time.Time, to time.Time) []domain.Ledgers {
	var ls []domain.Ledgers
	db := lR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&ls)
	return ls
}

func (lR ledgersRepository) Delete(id string) error {
	return lR.DB.Where("id = ?", id).Delete(&domain.Ledgers{}).Error
}
