package router

import "net/http"

// Context stores the context of a request.
type Context struct {
	Params         map[string]string
	Data           map[string]interface{}
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// NewContext returns a new Context object.
func NewContext(rw http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResponseWriter: rw,
		Data:           make(map[string]interface{}),
		Request:        r,
	}
}
