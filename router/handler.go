package router

import (
	"net/http"
)

// HandlerInterface is a Interface.
type HandlerInterface interface {
	Get(func(*Context))
	Post(func(*Context))
	Put(func(*Context))
	Patch(func(*Context))
	Delete(func(*Context))
	DoGet(*Context)
	DoPost(*Context)
	DoPut(*Context)
	DoPatch(*Context)
	DoDelete(*Context)
}

// Handler allows a type that can handle proper REST requests
type Handler struct {
	OnGet    func(context *Context)
	OnPost   func(context *Context)
	OnPut    func(context *Context)
	OnPatch  func(context *Context)
	OnDelete func(context *Context)
}

// NewHandler creates a new Handler for the Trie
func NewHandler() *Handler {
	return &Handler{
		OnGet:    nil,
		OnPost:   nil,
		OnPut:    nil,
		OnPatch:  nil,
		OnDelete: nil,
	}
}

// Get updates OnGet method.
func (handler *Handler) Get(handleFunc func(*Context)) {
	handler.OnGet = handleFunc
}

// Post updates OnPost method.
func (handler *Handler) Post(handleFunc func(*Context)) {
	handler.OnPost = handleFunc
}

// Put updates OnPut method.
func (handler *Handler) Put(handleFunc func(*Context)) {
	handler.OnPut = handleFunc
}

// Patch updates OnPatch method.
func (handler *Handler) Patch(handleFunc func(*Context)) {
	handler.OnPatch = handleFunc
}

// Delete updates OnDelete method.
func (handler *Handler) Delete(handleFunc func(*Context)) {
	handler.OnDelete = handleFunc
}

// DoNotFound handles when the request is not available
func (handler *Handler) DoNotFound(context *Context) {
	status := http.StatusMethodNotAllowed
	text := http.StatusText(status)
	context.ResponseWriter.WriteHeader(status)
	context.ResponseWriter.Write([]byte(text))
}

// DoGet handles running a handler's GET requests
func (handler *Handler) DoGet(context *Context) {
	if handler.OnGet == nil {
		handler.DoNotFound(context)
		return
	}
	handler.OnGet(context)
}

// DoPost handles running a handler's POST requests
func (handler *Handler) DoPost(context *Context) {
	if handler.OnPost == nil {
		handler.DoNotFound(context)
		return
	}
	handler.OnPost(context)
}

// DoPut handles running a handler's PUT requests
func (handler *Handler) DoPut(context *Context) {
	if handler.OnPut == nil {
		handler.DoNotFound(context)
		return
	}
	handler.OnPut(context)
}

// DoPatch handles running a handler's PATCH requests
func (handler *Handler) DoPatch(context *Context) {
	if handler.OnPatch == nil {
		handler.DoNotFound(context)
		return
	}
	handler.OnPatch(context)
}

// DoDelete handles running a handler's DELETE requests
func (handler *Handler) DoDelete(context *Context) {
	if handler.OnDelete == nil {
		handler.DoNotFound(context)
		return
	}
	handler.OnDelete(context)
}
