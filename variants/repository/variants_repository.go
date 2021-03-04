package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type variantsRepository struct {
	DB *gorm.DB
}

func NewVariantsRepository(db *gorm.DB) domain.VariantsRepository {
	return &variantsRepository{
		DB: db,
	}
}

func (vR variantsRepository) Save(v *domain.Variants) error {
	err := vR.DB.Save(v).Error
	return err
}

func (vR variantsRepository) Update(v *domain.Variants) error {
	err := vR.DB.Save(v).Error
	return err
}

func (vR variantsRepository) ByID(id string) (domain.Variants, error) {
	var v domain.Variants
	err := vR.DB.Where("id = ?", id).First(&v).Error
	return v, err
}
func (vR variantsRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Variants, lib.Paginator) {
	var vs []domain.Variants
	db := vR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&vs, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return vs, paginate
}

func (vR variantsRepository) FindRange(from time.Time, to time.Time) []domain.Variants {
	var vs []domain.Variants
	db := vR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&vs)
	return vs
}

func (vR variantsRepository) Delete(id string) error {
	return vR.DB.Where("id = ?", id).Delete(&domain.Variants{}).Error
}
