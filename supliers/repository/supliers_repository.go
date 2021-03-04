package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type supliersRepository struct {
	DB *gorm.DB
}

func NewSupliersRepository(db *gorm.DB) domain.SupliersRepository {
	return &supliersRepository{
		DB: db,
	}
}

func (sR supliersRepository) Save(s *domain.Supliers) error {
	err := sR.DB.Save(s).Error
	return err
}

func (sR supliersRepository) Update(s *domain.Supliers) error {
	err := sR.DB.Save(s).Error
	return err
}

func (sR supliersRepository) ByID(id string) (domain.Supliers, error) {
	var s domain.Supliers
	err := sR.DB.Where("id = ?", id).First(&s).Error
	return s, err
}
func (sR supliersRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Supliers, lib.Paginator) {
	var ss []domain.Supliers
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

func (sR supliersRepository) FindRange(from time.Time, to time.Time) []domain.Supliers {
	var ss []domain.Supliers
	db := sR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&ss)
	return ss
}

func (sR supliersRepository) Delete(id string) error {
	return sR.DB.Where("id = ?", id).Delete(&domain.Supliers{}).Error
}
