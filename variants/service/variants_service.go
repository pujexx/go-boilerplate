package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type variantsService struct {
	Repo domain.VariantsRepository
}

func NewVariantsService(repo domain.VariantsRepository) domain.VariantsService {
	return &variantsService{
		Repo: repo,
	}
}

func (vS variantsService) Save(v *domain.Variants) error {
	err := vS.Repo.Save(v)
	return err
}

func (vS variantsService) Update(v *domain.Variants) error {
	err := vS.Repo.Update(v)
	return err
}

func (vS variantsService) ByID(id string) (domain.Variants, error) {
	v, err := vS.Repo.ByID(id)
	return v, err
}
func (vS variantsService) Find(page int, filters []lib.FilterDomain) ([]domain.Variants, lib.Paginator) {
	vs, paginate := vS.Repo.Find(page, filters)
	return vs, paginate
}

func (vS variantsService) FindRange(from time.Time, to time.Time) (error, *[]domain.Variants) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	vs := vS.Repo.FindRange(from, to)
	return nil, &vs
}

func (vS variantsService) Delete(id string) error {
	return vS.Repo.Delete(id)
}
