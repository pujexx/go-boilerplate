package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type supliersService struct {
	Repo domain.SupliersRepository
}

func NewSupliersService(repo domain.SupliersRepository) domain.SupliersService {
	return &supliersService{
		Repo: repo,
	}
}

func (sS supliersService) Save(s *domain.Supliers) error {
	err := sS.Repo.Save(s)
	return err
}

func (sS supliersService) Update(s *domain.Supliers) error {
	err := sS.Repo.Update(s)
	return err
}

func (sS supliersService) ByID(id string) (domain.Supliers, error) {
	s, err := sS.Repo.ByID(id)
	return s, err
}
func (sS supliersService) Find(page int, filters []lib.FilterDomain) ([]domain.Supliers, lib.Paginator) {
	ss, paginate := sS.Repo.Find(page, filters)
	return ss, paginate
}

func (sS supliersService) FindRange(from time.Time, to time.Time) (error, *[]domain.Supliers) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	ss := sS.Repo.FindRange(from, to)
	return nil, &ss
}

func (sS supliersService) Delete(id string) error {
	return sS.Repo.Delete(id)
}
