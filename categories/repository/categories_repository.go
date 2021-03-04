package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type categoriesRepository struct {
	DB *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) domain.CategoriesRepository {
	return &categoriesRepository{
		DB: db,
	}
}

func (cR categoriesRepository) Save(c *domain.Categories) error {
	err := cR.DB.Save(c).Error
	return err
}

func (cR categoriesRepository) Update(c *domain.Categories) error {
	err := cR.DB.Save(c).Error
	return err
}

func (cR categoriesRepository) ByID(id string) (domain.Categories, error) {
	var c domain.Categories
	err := cR.DB.Where("id = ?", id).First(&c).Error
	return c, err
}
func (cR categoriesRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Categories, lib.Paginator) {
	var cs []domain.Categories
	db := cR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&cs, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return cs, paginate
}

func (cR categoriesRepository) FindRange(from time.Time, to time.Time) []domain.Categories {
	var cs []domain.Categories
	db := cR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&cs)
	return cs
}

func (cR categoriesRepository) Delete(id string) error {
	return cR.DB.Where("id = ?", id).Delete(&domain.Categories{}).Error
}
