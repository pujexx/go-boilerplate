package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type usersService struct {
	Repo domain.UsersRepository
}

func NewUsersService(repo domain.UsersRepository) domain.UsersService {
	return &usersService{
		Repo: repo,
	}
}

func (uS usersService) Save(u *domain.Users) error {
	err := uS.Repo.Save(u)
	return err
}

func (uS usersService) Update(u *domain.Users) error {
	err := uS.Repo.Update(u)
	return err
}

func (uS usersService) ByID(id string) (domain.Users, error) {
	u, err := uS.Repo.ByID(id)
	return u, err
}
func (uS usersService) Find(page int, filters []lib.FilterDomain) ([]domain.Users, lib.Paginator) {
	us, paginate := uS.Repo.Find(page, filters)
	return us, paginate
}

func (uS usersService) FindRange(from time.Time, to time.Time) (error, *[]domain.Users) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	us := uS.Repo.FindRange(from, to)
	return nil, &us
}

func (uS usersService) Delete(id string) error {
	return uS.Repo.Delete(id)
}
