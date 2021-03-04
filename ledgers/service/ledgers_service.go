package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type ledgersService struct {
	Repo domain.LedgersRepository
}

func NewLedgersService(repo domain.LedgersRepository) domain.LedgersService {
	return &ledgersService{
		Repo: repo,
	}
}

func (lS ledgersService) Save(l *domain.Ledgers) error {
	err := lS.Repo.Save(l)
	return err
}

func (lS ledgersService) Update(l *domain.Ledgers) error {
	err := lS.Repo.Update(l)
	return err
}

func (lS ledgersService) ByID(id string) (domain.Ledgers, error) {
	l, err := lS.Repo.ByID(id)
	return l, err
}
func (lS ledgersService) Find(page int, filters []lib.FilterDomain) ([]domain.Ledgers, lib.Paginator) {
	ls, paginate := lS.Repo.Find(page, filters)
	return ls, paginate
}

func (lS ledgersService) FindRange(from time.Time, to time.Time) (error, *[]domain.Ledgers) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	ls := lS.Repo.FindRange(from, to)
	return nil, &ls
}

func (lS ledgersService) Delete(id string) error {
	return lS.Repo.Delete(id)
}
