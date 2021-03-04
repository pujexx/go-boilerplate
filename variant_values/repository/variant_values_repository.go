package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type variantvaluesRepository struct {
	DB *gorm.DB
}

func NewVariantValuesRepository(db *gorm.DB) domain.VariantValuesRepository {
	return &variantvaluesRepository{
		DB: db,
	}
}

func (vvR variantvaluesRepository) Save(vv *domain.VariantValues) error {
	err := vvR.DB.Save(vv).Error
	return err
}

func (vvR variantvaluesRepository) Update(vv *domain.VariantValues) error {
	err := vvR.DB.Save(vv).Error
	return err
}

func (vvR variantvaluesRepository) ByID(id string) (domain.VariantValues, error) {
	var vv domain.VariantValues
	err := vvR.DB.Where("variant_options_id1 = ?", id).First(&vv).Error
	return vv, err
}
func (vvR variantvaluesRepository) Find(page int, filters []lib.FilterDomain) ([]domain.VariantValues, lib.Paginator) {
	var vvs []domain.VariantValues
	db := vvR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&vvs, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return vvs, paginate
}

func (vvR variantvaluesRepository) FindRange(from time.Time, to time.Time) []domain.VariantValues {
	var vvs []domain.VariantValues
	db := vvR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&vvs)
	return vvs
}

func (vvR variantvaluesRepository) Delete(id string) error {
	return vvR.DB.Where("variant_options_id1 = ?", id).Delete(&domain.VariantValues{}).Error
}
