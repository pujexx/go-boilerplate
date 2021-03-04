package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type productsService struct {
	Repo domain.ProductsRepository
}

func NewProductsService(repo domain.ProductsRepository) domain.ProductsService {
	return &productsService{
		Repo: repo,
	}
}

func (pS productsService) Save(p *domain.Products) error {
	err := pS.Repo.Save(p)
	return err
}

func (pS productsService) Update(p *domain.Products) error {
	err := pS.Repo.Update(p)
	return err
}

func (pS productsService) ByID(id string) (domain.Products, error) {
	p, err := pS.Repo.ByID(id)
	return p, err
}
func (pS productsService) Find(page int, filters []lib.FilterDomain) ([]domain.Products, lib.Paginator) {
	ps, paginate := pS.Repo.Find(page, filters)
	return ps, paginate
}

func (pS productsService) FindRange(from time.Time, to time.Time) (error, *[]domain.Products) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	ps := pS.Repo.FindRange(from, to)
	return nil, &ps
}

func (pS productsService) Delete(id string) error {
	return pS.Repo.Delete(id)
}
