package utils

import (
	"encoding/json"
	"github.com/cdsailing/pkg/log"
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

func (p *QueryParam) MapTo(out interface{}) {
	_bytes, err := json.Marshal(p)
	if err != nil {
		log.Warning("序列化数据失败: %s", err.Error())
		return
	}
	json.Unmarshal(_bytes, out)
}
