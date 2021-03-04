package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type rolesRepository struct {
	DB *gorm.DB
}

func NewRolesRepository(db *gorm.DB) domain.RolesRepository {
	return &rolesRepository{
		DB: db,
	}
}

func (rR rolesRepository) Save(r *domain.Roles) error {
	err := rR.DB.Save(r).Error
	return err
}

func (rR rolesRepository) Update(r *domain.Roles) error {
	err := rR.DB.Save(r).Error
	return err
}

func (rR rolesRepository) ByID(id string) (domain.Roles, error) {
	var r domain.Roles
	err := rR.DB.Where("id = ?", id).First(&r).Error
	return r, err
}
func (rR rolesRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Roles, lib.Paginator) {
	var rs []domain.Roles
	db := rR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&rs, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return rs, paginate
}

func (rR rolesRepository) FindRange(from time.Time, to time.Time) []domain.Roles {
	var rs []domain.Roles
	db := rR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&rs)
	return rs
}

func (rR rolesRepository) Delete(id string) error {
	return rR.DB.Where("id = ?", id).Delete(&domain.Roles{}).Error
}
