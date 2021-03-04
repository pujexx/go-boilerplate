package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type storesService struct {
	Repo domain.StoresRepository
}

func NewStoresService(repo domain.StoresRepository) domain.StoresService {
	return &storesService{
		Repo: repo,
	}
}

func (sS storesService) Save(s *domain.Stores) error {
	err := sS.Repo.Save(s)
	return err
}

func (sS storesService) Update(s *domain.Stores) error {
	err := sS.Repo.Update(s)
	return err
}

func (sS storesService) ByID(id string) (domain.Stores, error) {
	s, err := sS.Repo.ByID(id)
	return s, err
}
func (sS storesService) Find(page int, filters []lib.FilterDomain) ([]domain.Stores, lib.Paginator) {
	ss, paginate := sS.Repo.Find(page, filters)
	return ss, paginate
}

func (sS storesService) FindRange(from time.Time, to time.Time) (error, *[]domain.Stores) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	ss := sS.Repo.FindRange(from, to)
	return nil, &ss
}

func (sS storesService) Delete(id string) error {
	return sS.Repo.Delete(id)
}
