package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type ordersService struct {
	Repo domain.OrdersRepository
}

func NewOrdersService(repo domain.OrdersRepository) domain.OrdersService {
	return &ordersService{
		Repo: repo,
	}
}

func (oS ordersService) Save(o *domain.Orders) error {
	err := oS.Repo.Save(o)
	return err
}

func (oS ordersService) Update(o *domain.Orders) error {
	err := oS.Repo.Update(o)
	return err
}

func (oS ordersService) ByID(id string) (domain.Orders, error) {
	o, err := oS.Repo.ByID(id)
	return o, err
}
func (oS ordersService) Find(page int, filters []lib.FilterDomain) ([]domain.Orders, lib.Paginator) {
	os, paginate := oS.Repo.Find(page, filters)
	return os, paginate
}

func (oS ordersService) FindRange(from time.Time, to time.Time) (error, *[]domain.Orders) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	os := oS.Repo.FindRange(from, to)
	return nil, &os
}

func (oS ordersService) Delete(id string) error {
	return oS.Repo.Delete(id)
}
