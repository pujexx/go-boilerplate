package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type featuresService struct {
	Repo domain.FeaturesRepository
}

func NewFeaturesService(repo domain.FeaturesRepository) domain.FeaturesService {
	return &featuresService{
		Repo: repo,
	}
}

func (fS featuresService) Save(f *domain.Features) error {
	err := fS.Repo.Save(f)
	return err
}

func (fS featuresService) Update(f *domain.Features) error {
	err := fS.Repo.Update(f)
	return err
}

func (fS featuresService) ByID(id string) (domain.Features, error) {
	f, err := fS.Repo.ByID(id)
	return f, err
}
func (fS featuresService) Find(page int, filters []lib.FilterDomain) ([]domain.Features, lib.Paginator) {
	fs, paginate := fS.Repo.Find(page, filters)
	return fs, paginate
}

func (fS featuresService) FindRange(from time.Time, to time.Time) (error, *[]domain.Features) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	fs := fS.Repo.FindRange(from, to)
	return nil, &fs
}

func (fS featuresService) Delete(id string) error {
	return fS.Repo.Delete(id)
}
