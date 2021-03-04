package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type accessrolesService struct {
	Repo domain.AccessRolesRepository
}

func NewAccessRolesService(repo domain.AccessRolesRepository) domain.AccessRolesService {
	return &accessrolesService{
		Repo: repo,
	}
}

func (arS accessrolesService) Save(ar *domain.AccessRoles) error {
	err := arS.Repo.Save(ar)
	return err
}

func (arS accessrolesService) Update(ar *domain.AccessRoles) error {
	err := arS.Repo.Update(ar)
	return err
}

func (arS accessrolesService) ByID(id string) (domain.AccessRoles, error) {
	ar, err := arS.Repo.ByID(id)
	return ar, err
}
func (arS accessrolesService) Find(page int, filters []lib.FilterDomain) ([]domain.AccessRoles, lib.Paginator) {
	ars, paginate := arS.Repo.Find(page, filters)
	return ars, paginate
}

func (arS accessrolesService) FindRange(from time.Time, to time.Time) (error, *[]domain.AccessRoles) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	ars := arS.Repo.FindRange(from, to)
	return nil, &ars
}

func (arS accessrolesService) Delete(id string) error {
	return arS.Repo.Delete(id)
}
