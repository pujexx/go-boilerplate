package lib

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseDomain struct {
	ID        string     `gorm:"unique_index;primary_key;" json:"id,omitempty" validate:"omitempty,uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (baseRepository *BaseDomain) BeforeCreate(scope *gorm.DB) error {
	baseRepository.ID = uuid.New().String()
	return nil
}

type FilterDomain struct {
	Field string `json:"field"`
	Op string `json:"op"`
	Value interface{} `json:"value"`
}