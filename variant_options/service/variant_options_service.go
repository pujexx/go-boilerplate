package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type variantoptionsService struct {
	Repo domain.VariantOptionsRepository
}

func NewVariantOptionsService(repo domain.VariantOptionsRepository) domain.VariantOptionsService {
	return &variantoptionsService{
		Repo: repo,
	}
}

func (voS variantoptionsService) Save(vo *domain.VariantOptions) error {
	err := voS.Repo.Save(vo)
	return err
}

func (voS variantoptionsService) Update(vo *domain.VariantOptions) error {
	err := voS.Repo.Update(vo)
	return err
}

func (voS variantoptionsService) ByID(id string) (domain.VariantOptions, error) {
	vo, err := voS.Repo.ByID(id)
	return vo, err
}
func (voS variantoptionsService) Find(page int, filters []lib.FilterDomain) ([]domain.VariantOptions, lib.Paginator) {
	vos, paginate := voS.Repo.Find(page, filters)
	return vos, paginate
}

func (voS variantoptionsService) FindRange(from time.Time, to time.Time) (error, *[]domain.VariantOptions) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	vos := voS.Repo.FindRange(from, to)
	return nil, &vos
}

func (voS variantoptionsService) Delete(id string) error {
	return voS.Repo.Delete(id)
}
