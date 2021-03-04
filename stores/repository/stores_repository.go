package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type storesRepository struct {
	DB *gorm.DB
}

func NewStoresRepository(db *gorm.DB) domain.StoresRepository {
	return &storesRepository{
		DB: db,
	}
}

func (sR storesRepository) Save(s *domain.Stores) error {
	err := sR.DB.Save(s).Error
	return err
}

func (sR storesRepository) Update(s *domain.Stores) error {
	err := sR.DB.Save(s).Error
	return err
}

func (sR storesRepository) ByID(id string) (domain.Stores, error) {
	var s domain.Stores
	err := sR.DB.Where("id = ?", id).First(&s).Error
	return s, err
}
func (sR storesRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Stores, lib.Paginator) {
	var ss []domain.Stores
	db := sR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&ss, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return ss, paginate
}

func (sR storesRepository) FindRange(from time.Time, to time.Time) []domain.Stores {
	var ss []domain.Stores
	db := sR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&ss)
	return ss
}

func (sR storesRepository) Delete(id string) error {
	return sR.DB.Where("id = ?", id).Delete(&domain.Stores{}).Error
}
