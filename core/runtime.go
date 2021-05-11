package core

import (
	"bytes"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"runtime"
)

func LogStackOnRecover(panicReason interface{}, httpWriter http.ResponseWriter) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("recover from panic situation: - %v\r\n", panicReason))
	for i := 2; ; i += 1 {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", file, line))
	}
	fmt.Errorf(buffer.String())
	httpWriter.WriteHeader(http.StatusInternalServerError)
	httpWriter.Write([]byte("recover from panic situation"))
}

var Container = restful.DefaultContainer

type ContainerBuilder []func(c *restful.Container) error

const MimeMergePatchJson = "application/merge-patch+json"
const MimeJsonPatchJson = "application/json-patch+json"

func init() {
	restful.RegisterEntityAccessor(MimeMergePatchJson, restful.NewEntityAccessorJSON(restful.MIME_JSON))
	restful.RegisterEntityAccessor(MimeJsonPatchJson, restful.NewEntityAccessorJSON(restful.MIME_JSON))
}

func NewWebService(group string, ApiPrefix string) *restful.WebService {
	service := restful.WebService{}
	service.Path(ApiPrefix + "/" + group)
	return &service
}
func (builder *ContainerBuilder) AddToContainer(c *restful.Container) error {
	for _, f := range *builder {
		if err := f(c); err != nil {
			return err
		}
	}
	return nil
}

func (builder *ContainerBuilder) Register(funcs ...func(*restful.Container) error) {
	for _, f := range funcs {
		*builder = append(*builder, f)
	}
}

func NewContainerBuilder(funcs ...func(*restful.Container) error) ContainerBuilder {
	var builder ContainerBuilder
	builder.Register(funcs...)
	return builder
}
