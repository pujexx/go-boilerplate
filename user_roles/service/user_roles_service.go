package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type userrolesService struct {
	Repo domain.UserRolesRepository
}

func NewUserRolesService(repo domain.UserRolesRepository) domain.UserRolesService {
	return &userrolesService{
		Repo: repo,
	}
}

func (urS userrolesService) Save(ur *domain.UserRoles) error {
	err := urS.Repo.Save(ur)
	return err
}

func (urS userrolesService) Update(ur *domain.UserRoles) error {
	err := urS.Repo.Update(ur)
	return err
}

func (urS userrolesService) ByID(id string) (domain.UserRoles, error) {
	ur, err := urS.Repo.ByID(id)
	return ur, err
}
func (urS userrolesService) Find(page int, filters []lib.FilterDomain) ([]domain.UserRoles, lib.Paginator) {
	urs, paginate := urS.Repo.Find(page, filters)
	return urs, paginate
}

func (urS userrolesService) FindRange(from time.Time, to time.Time) (error, *[]domain.UserRoles) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	urs := urS.Repo.FindRange(from, to)
	return nil, &urs
}

func (urS userrolesService) Delete(id string) error {
	return urS.Repo.Delete(id)
}
