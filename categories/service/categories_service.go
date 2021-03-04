package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type categoriesService struct {
	Repo domain.CategoriesRepository
}

func NewCategoriesService(repo domain.CategoriesRepository) domain.CategoriesService {
	return &categoriesService{
		Repo: repo,
	}
}

func (cS categoriesService) Save(c *domain.Categories) error {
	err := cS.Repo.Save(c)
	return err
}

func (cS categoriesService) Update(c *domain.Categories) error {
	err := cS.Repo.Update(c)
	return err
}

func (cS categoriesService) ByID(id string) (domain.Categories, error) {
	c, err := cS.Repo.ByID(id)
	return c, err
}
func (cS categoriesService) Find(page int, filters []lib.FilterDomain) ([]domain.Categories, lib.Paginator) {
	cs, paginate := cS.Repo.Find(page, filters)
	return cs, paginate
}

func (cS categoriesService) FindRange(from time.Time, to time.Time) (error, *[]domain.Categories) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	cs := cS.Repo.FindRange(from, to)
	return nil, &cs
}

func (cS categoriesService) Delete(id string) error {
	return cS.Repo.Delete(id)
}
