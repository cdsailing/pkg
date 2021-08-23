package pager

import "reflect"

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
	PageSize int
	Page     int
	Keyword  string
	Order    string
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
	if pager.Total%int64(pager.Size) == 0 {
		pager.Count = int(pager.Total % int64(pager.Size))
	} else {
		pager.Count = int(pager.Total%int64(pager.Size)) + 1
	}
	return pager

	return nil
}
