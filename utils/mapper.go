package utils

import (
	"github.com/mitchellh/mapstructure"
	"net/url"
)

type QueryParam struct {
	url.Values
}

func Map(input interface{}, out interface{}) {
	mapstructure.Decode(input, out)
}

func (p *QueryParam) Map(out interface{}) {
	result := make(map[string]interface{})
	for s := range p.Values {
		result[s] = p.Values.Get(s)
	}
	Map(result, out)
}
