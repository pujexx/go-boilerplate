package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type transactionsService struct {
	Repo domain.TransactionsRepository
}

func NewTransactionsService(repo domain.TransactionsRepository) domain.TransactionsService {
	return &transactionsService{
		Repo: repo,
	}
}

func (tS transactionsService) Save(t *domain.Transactions) error {
	err := tS.Repo.Save(t)
	return err
}

func (tS transactionsService) Update(t *domain.Transactions) error {
	err := tS.Repo.Update(t)
	return err
}

func (tS transactionsService) ByID(id string) (domain.Transactions, error) {
	t, err := tS.Repo.ByID(id)
	return t, err
}
func (tS transactionsService) Find(page int, filters []lib.FilterDomain) ([]domain.Transactions, lib.Paginator) {
	ts, paginate := tS.Repo.Find(page, filters)
	return ts, paginate
}

func (tS transactionsService) FindRange(from time.Time, to time.Time) (error, *[]domain.Transactions) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	ts := tS.Repo.FindRange(from, to)
	return nil, &ts
}

func (tS transactionsService) Delete(id string) error {
	return tS.Repo.Delete(id)
}
