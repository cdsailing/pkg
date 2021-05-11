package pager

import "reflect"

type Pager struct {
	Items []interface{} `json:"items" description:"paging data"`
	Total int64         `json:"total" description:"total count"`
	Order string        `json:"order"`
	Page  int           `json:"page"`
	Count int           `json:"count"`
	Size  int           `json:"size"`
}

func ToPager(source interface{}, total int64) *Pager {
	pager := &Pager{Total: total}
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
	pager.Items = result
	return pager

	return nil
}