package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type extraordersService struct {
	Repo domain.ExtraOrdersRepository
}

func NewExtraOrdersService(repo domain.ExtraOrdersRepository) domain.ExtraOrdersService {
	return &extraordersService{
		Repo: repo,
	}
}

func (eoS extraordersService) Save(eo *domain.ExtraOrders) error {
	err := eoS.Repo.Save(eo)
	return err
}

func (eoS extraordersService) Update(eo *domain.ExtraOrders) error {
	err := eoS.Repo.Update(eo)
	return err
}

func (eoS extraordersService) ByID(id string) (domain.ExtraOrders, error) {
	eo, err := eoS.Repo.ByID(id)
	return eo, err
}
func (eoS extraordersService) Find(page int, filters []lib.FilterDomain) ([]domain.ExtraOrders, lib.Paginator) {
	eos, paginate := eoS.Repo.Find(page, filters)
	return eos, paginate
}

func (eoS extraordersService) FindRange(from time.Time, to time.Time) (error, *[]domain.ExtraOrders) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	eos := eoS.Repo.FindRange(from, to)
	return nil, &eos
}

func (eoS extraordersService) Delete(id string) error {
	return eoS.Repo.Delete(id)
}
