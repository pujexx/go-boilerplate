package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type membersRepository struct {
	DB *gorm.DB
}

func NewMembersRepository(db *gorm.DB) domain.MembersRepository {
	return &membersRepository{
		DB: db,
	}
}

func (mR membersRepository) Save(m *domain.Members) error {
	err := mR.DB.Save(m).Error
	return err
}

func (mR membersRepository) Update(m *domain.Members) error {
	err := mR.DB.Save(m).Error
	return err
}

func (mR membersRepository) ByID(id string) (domain.Members, error) {
	var m domain.Members
	err := mR.DB.Where("id = ?", id).First(&m).Error
	return m, err
}
func (mR membersRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Members, lib.Paginator) {
	var ms []domain.Members
	db := mR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&ms, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return ms, paginate
}

func (mR membersRepository) FindRange(from time.Time, to time.Time) []domain.Members {
	var ms []domain.Members
	db := mR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&ms)
	return ms
}

func (mR membersRepository) Delete(id string) error {
	return mR.DB.Where("id = ?", id).Delete(&domain.Members{}).Error
}
