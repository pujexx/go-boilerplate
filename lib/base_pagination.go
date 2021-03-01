package lib

import (
	"gorm.io/gorm"
	"math"
)

type Param struct {
	DB      *gorm.DB
	Page    int
	PerPage int
	OrderBy int
}

type Paginator struct {
	Data interface{}   `json:"data"`
	Meta PaginatorMeta `json:"meta"`
}

type PaginatorMeta struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	Total     int `json:"total"`
	PerPage   int `json:"per_page"`
}

func Paginate(result interface{}, param *Param) Paginator {
	DB := param.DB

	var count int64
	var offset int

	DB.Model(result).Count(&count)

	var paginator Paginator
	paginator.Meta.TotalPage = int(math.Ceil(float64(count) / float64(param.PerPage)))
	paginator.Meta.Page = param.Page
	paginator.Meta.Total = int(count)
	paginator.Meta.PerPage = param.PerPage
	if param.Page == 1 {
		offset = 0
	} else {
		offset = (param.Page - 1) * param.PerPage
	}
	DB.Limit(param.PerPage).Offset(offset).Find(result)
	paginator.Data = &result

	return paginator
}
