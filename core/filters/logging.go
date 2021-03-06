package filters

import (
	"github.com/cdsailing/pkg/log"
	"github.com/emicklei/go-restful/v3"
	"net"
	"strings"
	"time"
)

func Logging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	start := time.Now()
	chain.ProcessFilter(req, resp)
	log.Infof("%s - \"%s %s %s\" %d %d %dms",
		getRequestIP(req),
		req.Request.Method,
		req.Request.RequestURI,
		req.Request.Proto,
		resp.StatusCode(),
		resp.ContentLength(),
		time.Since(start)/time.Millisecond,
	)
}

func getRequestIP(req *restful.Request) string {
	address := strings.Trim(req.Request.Header.Get("X-Real-Ip"), " ")
	if address != "" {
		return address
	}

	address = strings.Trim(req.Request.Header.Get("X-Forwarded-For"), " ")
	if address != "" {
		return address
	}

	address, _, err := net.SplitHostPort(req.Request.RemoteAddr)
	if err != nil {
		return req.Request.RemoteAddr
	}

	return address
}
