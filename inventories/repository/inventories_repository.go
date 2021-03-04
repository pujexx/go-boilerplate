package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type inventoriesRepository struct {
	DB *gorm.DB
}

func NewInventoriesRepository(db *gorm.DB) domain.InventoriesRepository {
	return &inventoriesRepository{
		DB: db,
	}
}

func (iR inventoriesRepository) Save(i *domain.Inventories) error {
	err := iR.DB.Save(i).Error
	return err
}

func (iR inventoriesRepository) Update(i *domain.Inventories) error {
	err := iR.DB.Save(i).Error
	return err
}

func (iR inventoriesRepository) ByID(id string) (domain.Inventories, error) {
	var i domain.Inventories
	err := iR.DB.Where("id = ?", id).First(&i).Error
	return i, err
}
func (iR inventoriesRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Inventories, lib.Paginator) {
	var is []domain.Inventories
	db := iR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&is, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return is, paginate
}

func (iR inventoriesRepository) FindRange(from time.Time, to time.Time) []domain.Inventories {
	var is []domain.Inventories
	db := iR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&is)
	return is
}

func (iR inventoriesRepository) Delete(id string) error {
	return iR.DB.Where("id = ?", id).Delete(&domain.Inventories{}).Error
}
