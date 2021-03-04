package repository

import (
	"fmt"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"gorm.io/gorm"
	"strings"
	"time"
)

type usersRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) domain.UsersRepository {
	return &usersRepository{
		DB: db,
	}
}

func (uR usersRepository) Save(u *domain.Users) error {
	err := uR.DB.Save(u).Error
	return err
}

func (uR usersRepository) Update(u *domain.Users) error {
	err := uR.DB.Save(u).Error
	return err
}

func (uR usersRepository) ByID(id string) (domain.Users, error) {
	var u domain.Users
	err := uR.DB.Where("id = ?", id).First(&u).Error
	return u, err
}
func (uR usersRepository) Find(page int, filters []lib.FilterDomain) ([]domain.Users, lib.Paginator) {
	var us []domain.Users
	db := uR.DB
	conditions := []string{}

	for _, where := range filters {
		conditions = append(conditions, fmt.Sprintf("%v %v %v", where.Field, where.Op, where.Value))
	}
	if len(filters) > 0 {
		db = db.Where(strings.Join(conditions, " AND "))
	}

	paginate := lib.Paginate(&us, &lib.Param{
		DB:      db,
		Page:    page,
		PerPage: 20,
		OrderBy: 0,
	})
	return us, paginate
}

func (uR usersRepository) FindRange(from time.Time, to time.Time) []domain.Users {
	var us []domain.Users
	db := uR.DB
	db = db.Where("created_at <= ? and created_at >= ?", from, to)
	db.Find(&us)
	return us
}

func (uR usersRepository) Delete(id string) error {
	return uR.DB.Where("id = ?", id).Delete(&domain.Users{}).Error
}
