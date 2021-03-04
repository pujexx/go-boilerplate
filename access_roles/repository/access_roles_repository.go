package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type accessrolesRepository struct {
	DB *gorm.DB
}

func NewAccessRolesRepository(db *gorm.DB) domain.AccessRolesRepository {
	return &accessrolesRepository{
		DB: db,
	}
}

func (arR accessrolesRepository) Save(ar *domain.AccessRoles) error {
	err := arR.DB.Save(ar).Error
	return err
}

func (arR accessrolesRepository) Update(ar *domain.AccessRoles) error {
	err := arR.DB.Save(ar).Error
	return err
}

func (arR accessrolesRepository) ByID(id string) (domain.AccessRoles, error) {
	var ar domain.AccessRoles
	err := arR.DB.Where("id = ?", id).First(&ar).Error
	return ar, err
}
func (arR accessrolesRepository) Find(page int, filters []lib.FilterDomain) ([]domain.AccessRoles, lib.Paginator) {
	var ars []domain.AccessRoles
	db := arR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&ars, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return ars, paginate
}

func (arR accessrolesRepository) FindRange(from time.Time, to time.Time) []domain.AccessRoles {
	var ars []domain.AccessRoles
	db := arR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&ars)
	return ars
}

func (arR accessrolesRepository) Delete(id string) error {
	return arR.DB.Where("id = ?", id).Delete(&domain.AccessRoles{}).Error
}
