package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type featuresRepository struct {
	DB *gorm.DB
}

func NewFeaturesRepository(db *gorm.DB) domain.FeaturesRepository {
	return &featuresRepository{
		DB: db,
	}
}

func (fR featuresRepository) Save(f *domain.Features) error {
	err := fR.DB.Save(f).Error
	return err
}

func (fR featuresRepository) Update(f *domain.Features) error {
	err := fR.DB.Save(f).Error
	return err
}

func (fR featuresRepository) ByID(id string) (domain.Features, error) {
	var f domain.Features
	err := fR.DB.Where("id = ?", id).First(&f).Error
	return f, err
}
func (fR featuresRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Features, lib.Paginator) {
	var fs []domain.Features
	db := fR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&fs, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return fs, paginate
}

func (fR featuresRepository) FindRange(from time.Time, to time.Time) []domain.Features {
	var fs []domain.Features
	db := fR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&fs)
	return fs
}

func (fR featuresRepository) Delete(id string) error {
	return fR.DB.Where("id = ?", id).Delete(&domain.Features{}).Error
}
