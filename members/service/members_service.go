package service

import (
	"errors"
	"github.com/pujexx/go-boilerplate/domain"
	"github.com/pujexx/go-boilerplate/lib"
	"time"
)

type membersService struct {
	Repo domain.MembersRepository
}

func NewMembersService(repo domain.MembersRepository) domain.MembersService {
	return &membersService{
		Repo: repo,
	}
}

func (mS membersService) Save(m *domain.Members) error {
	err := mS.Repo.Save(m)
	return err
}

func (mS membersService) Update(m *domain.Members) error {
	err := mS.Repo.Update(m)
	return err
}

func (mS membersService) ByID(id string) (domain.Members, error) {
	m, err := mS.Repo.ByID(id)
	return m, err
}
func (mS membersService) Find(page int, filters []lib.FilterDomain) ([]domain.Members, lib.Paginator) {
	ms, paginate := mS.Repo.Find(page, filters)
	return ms, paginate
}

func (mS membersService) FindRange(from time.Time, to time.Time) (error, *[]domain.Members) {
	dateRange := lib.DateRange(from, to)
	if dateRange > 30 || dateRange < 0 {
		return errors.New("date out of range"), nil
	}
	ms := mS.Repo.FindRange(from, to)
	return nil, &ms
}

func (mS membersService) Delete(id string) error {
	return mS.Repo.Delete(id)
}
