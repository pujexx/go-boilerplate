package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type productsRepository struct {
	DB *gorm.DB
}

func NewProductsRepository(db *gorm.DB) domain.ProductsRepository {
	return &productsRepository{
		DB: db,
	}
}

func (pR productsRepository) Save(p *domain.Products) error {
	err := pR.DB.Save(p).Error
	return err
}

func (pR productsRepository) Update(p *domain.Products) error {
	err := pR.DB.Save(p).Error
	return err
}

func (pR productsRepository) ByID(id string) (domain.Products, error) {
	var p domain.Products
	err := pR.DB.Where("id = ?", id).First(&p).Error
	return p, err
}
func (pR productsRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Products, lib.Paginator) {
	var ps []domain.Products
	db := pR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&ps, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return ps, paginate
}

func (pR productsRepository) FindRange(from time.Time, to time.Time) []domain.Products {
	var ps []domain.Products
	db := pR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&ps)
	return ps
}

func (pR productsRepository) Delete(id string) error {
	return pR.DB.Where("id = ?", id).Delete(&domain.Products{}).Error
}
