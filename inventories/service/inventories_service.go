package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type inventoriesService struct {
	Repo domain.InventoriesRepository
}

func NewInventoriesService(repo domain.InventoriesRepository) domain.InventoriesService {
	return &inventoriesService{
		Repo: repo,
	}
}

func (iS inventoriesService) Save(i *domain.Inventories) error {
	err := iS.Repo.Save(i)
	return err
}

func (iS inventoriesService) Update(i *domain.Inventories) error {
	err := iS.Repo.Update(i)
	return err
}

func (iS inventoriesService) ByID(id string) (domain.Inventories, error) {
	i, err := iS.Repo.ByID(id)
	return i, err
}
func (iS inventoriesService) Find(page int, filters []lib.FilterDomain) ([]domain.Inventories, lib.Paginator) {
	is, paginate := iS.Repo.Find(page, filters)
	return is, paginate
}

func (iS inventoriesService) FindRange(from time.Time, to time.Time) (error, *[]domain.Inventories) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	is := iS.Repo.FindRange(from, to)
	return nil, &is
}

func (iS inventoriesService) Delete(id string) error {
	return iS.Repo.Delete(id)
}
