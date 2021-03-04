package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type rolesService struct {
	Repo domain.RolesRepository
}

func NewRolesService(repo domain.RolesRepository) domain.RolesService {
	return &rolesService{
		Repo: repo,
	}
}

func (rS rolesService) Save(r *domain.Roles) error {
	err := rS.Repo.Save(r)
	return err
}

func (rS rolesService) Update(r *domain.Roles) error {
	err := rS.Repo.Update(r)
	return err
}

func (rS rolesService) ByID(id string) (domain.Roles, error) {
	r, err := rS.Repo.ByID(id)
	return r, err
}
func (rS rolesService) Find(page int, filters []lib.FilterDomain) ([]domain.Roles, lib.Paginator) {
	rs, paginate := rS.Repo.Find(page, filters)
	return rs, paginate
}

func (rS rolesService) FindRange(from time.Time, to time.Time) (error, *[]domain.Roles) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	rs := rS.Repo.FindRange(from, to)
	return nil, &rs
}

func (rS rolesService) Delete(id string) error {
	return rS.Repo.Delete(id)
}
