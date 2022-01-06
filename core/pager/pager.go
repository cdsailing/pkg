package pager

import (
	"github.com/cdsailing/pkg/config"
	"reflect"
)

type PageResult struct {
	Data  []interface{} `json:"data" description:"paging data"`
	Total int64         `json:"totalCount" description:"total count"`
	Page  int           `json:"page"`
	Count int           `json:"pageCount"`
	Size  int           `json:"pageSize"`
}

/**
分页查询
*/
type PageQuery struct {
	PageSize int    `json:"pageSize" form:"pageSize"`
	Page     int    `json:"page" form:"page"`
	Keyword  string `json:"keyword" form:"keyword"`
	Order    string `json:"order" form:"order"`
}

func (p *PageQuery) GetPager() {
	if p.PageSize <= 0 {
		if config.Conf != nil && config.Conf.Server.PageSize > 0 {
			p.PageSize = config.Conf.Server.PageSize
		} else {
			p.PageSize = 20
		}
	}
	if p.Page <= 0 {
		p.Page = 1
	}
}

func ToPager(source interface{}, total int64, query PageQuery) *PageResult {
	pager := &PageResult{Total: total}
	temp := make([]interface{}, 0)
	switch reflect.TypeOf(source).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(source)
		for i := 0; i < s.Len(); i++ {
			t := s.Index(i)
			temp = append(temp, t.Interface())
		}
	}
	result := make([]interface{}, 0)
	for _, a := range temp {
		result = append(result, a)
	}
	pager.Data = result
	pager.Page = query.Page
	pager.Size = query.PageSize
	if pager.Total/int64(pager.Size) == 0 {
		pager.Count = int(pager.Total / int64(pager.Size))
	} else {
		pager.Count = int(pager.Total/int64(pager.Size)) + 1
	}
	return pager

	return nil
}
