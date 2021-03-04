package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type variantvaluesService struct {
	Repo domain.VariantValuesRepository
}

func NewVariantValuesService(repo domain.VariantValuesRepository) domain.VariantValuesService {
	return &variantvaluesService{
		Repo: repo,
	}
}

func (vvS variantvaluesService) Save(vv *domain.VariantValues) error {
	err := vvS.Repo.Save(vv)
	return err
}

func (vvS variantvaluesService) Update(vv *domain.VariantValues) error {
	err := vvS.Repo.Update(vv)
	return err
}

func (vvS variantvaluesService) ByID(id string) (domain.VariantValues, error) {
	vv, err := vvS.Repo.ByID(id)
	return vv, err
}
func (vvS variantvaluesService) Find(page int, filters []lib.FilterDomain) ([]domain.VariantValues, lib.Paginator) {
	vvs, paginate := vvS.Repo.Find(page, filters)
	return vvs, paginate
}

func (vvS variantvaluesService) FindRange(from time.Time, to time.Time) (error, *[]domain.VariantValues) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	vvs := vvS.Repo.FindRange(from, to)
	return nil, &vvs
}

func (vvS variantvaluesService) Delete(id string) error {
	return vvS.Repo.Delete(id)
}
