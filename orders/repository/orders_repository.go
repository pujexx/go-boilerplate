package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ordersRepository struct {
	DB *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) domain.OrdersRepository {
	return &ordersRepository{
		DB: db,
	}
}

func (oR ordersRepository) Save(o *domain.Orders) error {
	err := oR.DB.Save(o).Error
	return err
}

func (oR ordersRepository) Update(o *domain.Orders) error {
	err := oR.DB.Save(o).Error
	return err
}

func (oR ordersRepository) ByID(id string) (domain.Orders, error) {
	var o domain.Orders
	err := oR.DB.Where("id = ?", id).First(&o).Error
	return o, err
}
func (oR ordersRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Orders, lib.Paginator) {
	var os []domain.Orders
	db := oR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&os, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return os, paginate
}

func (oR ordersRepository) FindRange(from time.Time, to time.Time) []domain.Orders {
	var os []domain.Orders
	db := oR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&os)
	return os
}

func (oR ordersRepository) Delete(id string) error {
	return oR.DB.Where("id = ?", id).Delete(&domain.Orders{}).Error
}
