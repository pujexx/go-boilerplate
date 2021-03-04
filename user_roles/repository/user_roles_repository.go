package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type userrolesRepository struct {
	DB *gorm.DB
}

func NewUserRolesRepository(db *gorm.DB) domain.UserRolesRepository {
	return &userrolesRepository{
		DB: db,
	}
}

func (urR userrolesRepository) Save(ur *domain.UserRoles) error {
	err := urR.DB.Save(ur).Error
	return err
}

func (urR userrolesRepository) Update(ur *domain.UserRoles) error {
	err := urR.DB.Save(ur).Error
	return err
}

func (urR userrolesRepository) ByID(id string) (domain.UserRoles, error) {
	var ur domain.UserRoles
	err := urR.DB.Where("store_id = ?", id).First(&ur).Error
	return ur, err
}
func (urR userrolesRepository) Find(page int, filters []lib.FilterDomain) ([]domain.UserRoles, lib.Paginator) {
	var urs []domain.UserRoles
	db := urR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&urs, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return urs, paginate
}

func (urR userrolesRepository) FindRange(from time.Time, to time.Time) []domain.UserRoles {
	var urs []domain.UserRoles
	db := urR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&urs)
	return urs
}

func (urR userrolesRepository) Delete(id string) error {
	return urR.DB.Where("store_id = ?", id).Delete(&domain.UserRoles{}).Error
}
