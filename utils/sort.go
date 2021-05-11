package utils

import (
	"reflect"
	"sort"
	"time"
)

type Items struct {
	data  interface{}
	field string
}

func (items *Items) Len() int {
	if reflect.ValueOf(items.data).Kind() != reflect.Slice {
		return -1
	}
	return reflect.ValueOf(items.data).Len()
}

func (items *Items) Less(i, j int) bool {
	a := reflect.ValueOf(items.data).Index(i)
	b := reflect.ValueOf(items.data).Index(j)
	if a.Kind() == reflect.Ptr {
		a = a.Elem()
	}
	if b.Kind() == reflect.Ptr {
		b = b.Elem()
	}

	va, _ := a.FieldByName(items.field).Interface().(time.Time)
	vb, _ := b.FieldByName(items.field).Interface().(time.Time)
	return va.Before(vb)
}

func (items *Items) Swap(i, j int) {
	reflect.Swapper(items.data)(i, j)
}

func SortItems(i interface{}, str string) {
	if reflect.ValueOf(i).Kind() != reflect.Slice {
		return
	}
	a := &Items{
		data:  i,
		field: str,
	}
	sort.Sort(a)
}
