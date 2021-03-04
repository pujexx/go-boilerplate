package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type cdnsService struct {
	Repo domain.CdnsRepository
}

func NewCdnsService(repo domain.CdnsRepository) domain.CdnsService {
	return &cdnsService{
		Repo: repo,
	}
}

func (cS cdnsService) Save(c *domain.Cdns) error {
	err := cS.Repo.Save(c)
	return err
}

func (cS cdnsService) Update(c *domain.Cdns) error {
	err := cS.Repo.Update(c)
	return err
}

func (cS cdnsService) ByID(id string) (domain.Cdns, error) {
	c, err := cS.Repo.ByID(id)
	return c, err
}
func (cS cdnsService) Find(page int, filters []lib.FilterDomain) ([]domain.Cdns, lib.Paginator) {
	cs, paginate := cS.Repo.Find(page, filters)
	return cs, paginate
}

func (cS cdnsService) FindRange(from time.Time, to time.Time) (error, *[]domain.Cdns) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	cs := cS.Repo.FindRange(from, to)
	return nil, &cs
}

func (cS cdnsService) Delete(id string) error {
	return cS.Repo.Delete(id)
}
