package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type variantoptionsRepository struct {
	DB *gorm.DB
}

func NewVariantOptionsRepository(db *gorm.DB) domain.VariantOptionsRepository {
	return &variantoptionsRepository{
		DB: db,
	}
}

func (voR variantoptionsRepository) Save(vo *domain.VariantOptions) error {
	err := voR.DB.Save(vo).Error
	return err
}

func (voR variantoptionsRepository) Update(vo *domain.VariantOptions) error {
	err := voR.DB.Save(vo).Error
	return err
}

func (voR variantoptionsRepository) ByID(id string) (domain.VariantOptions, error) {
	var vo domain.VariantOptions
	err := voR.DB.Where("id = ?", id).First(&vo).Error
	return vo, err
}
func (voR variantoptionsRepository) Find(page int, filters []lib.FilterDomain) ([]domain.VariantOptions, lib.Paginator) {
	var vos []domain.VariantOptions
	db := voR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&vos, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return vos, paginate
}

func (voR variantoptionsRepository) FindRange(from time.Time, to time.Time) []domain.VariantOptions {
	var vos []domain.VariantOptions
	db := voR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&vos)
	return vos
}

func (voR variantoptionsRepository) Delete(id string) error {
	return voR.DB.Where("id = ?", id).Delete(&domain.VariantOptions{}).Error
}
