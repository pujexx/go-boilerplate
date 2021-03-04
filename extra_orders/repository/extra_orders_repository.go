package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type extraordersRepository struct {
	DB *gorm.DB
}

func NewExtraOrdersRepository(db *gorm.DB) domain.ExtraOrdersRepository {
	return &extraordersRepository{
		DB: db,
	}
}

func (eoR extraordersRepository) Save(eo *domain.ExtraOrders) error {
	err := eoR.DB.Save(eo).Error
	return err
}

func (eoR extraordersRepository) Update(eo *domain.ExtraOrders) error {
	err := eoR.DB.Save(eo).Error
	return err
}

func (eoR extraordersRepository) ByID(id string) (domain.ExtraOrders, error) {
	var eo domain.ExtraOrders
	err := eoR.DB.Where("id = ?", id).First(&eo).Error
	return eo, err
}
func (eoR extraordersRepository) Find(page int, filters []lib.FilterDomain) ([]domain.ExtraOrders, lib.Paginator) {
	var eos []domain.ExtraOrders
	db := eoR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&eos, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return eos, paginate
}

func (eoR extraordersRepository) FindRange(from time.Time, to time.Time) []domain.ExtraOrders {
	var eos []domain.ExtraOrders
	db := eoR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&eos)
	return eos
}

func (eoR extraordersRepository) Delete(id string) error {
	return eoR.DB.Where("id = ?", id).Delete(&domain.ExtraOrders{}).Error
}
