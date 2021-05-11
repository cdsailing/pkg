package response

import (
	"github.com/cdsailing/pkg/log"
	"github.com/emicklei/go-restful/v3"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

//Success
func Success(response *restful.Response, data interface{}) {
	r := Result{1, "success", data}
	response.WriteEntity(r)
}

//Fail
func Fail(response *restful.Response, error error) {
	r := Result{0, error.Error(), nil}
	log.Error(error.Error)
	response.WriteEntity(r)
}
