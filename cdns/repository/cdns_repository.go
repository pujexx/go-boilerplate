package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type cdnsRepository struct {
	DB *gorm.DB
}

func NewCdnsRepository(db *gorm.DB) domain.CdnsRepository {
	return &cdnsRepository{
		DB: db,
	}
}

func (cR cdnsRepository) Save(c *domain.Cdns) error {
	err := cR.DB.Save(c).Error
	return err
}

func (cR cdnsRepository) Update(c *domain.Cdns) error {
	err := cR.DB.Save(c).Error
	return err
}

func (cR cdnsRepository) ByID(id string) (domain.Cdns, error) {
	var c domain.Cdns
	err := cR.DB.Where("id = ?", id).First(&c).Error
	return c, err
}
func (cR cdnsRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Cdns, lib.Paginator) {
	var cs []domain.Cdns
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

func (cR cdnsRepository) FindRange(from time.Time, to time.Time) []domain.Cdns {
	var cs []domain.Cdns
	db := cR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&cs)
	return cs
}

func (cR cdnsRepository) Delete(id string) error {
	return cR.DB.Where("id = ?", id).Delete(&domain.Cdns{}).Error
}
