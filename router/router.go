package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/victor-cabrera/fe-calc/server/consts"
	"github.com/victor-cabrera/fe-calc/server/logger"
	"github.com/victor-cabrera/fe-calc/server/types"
)

// Router is our Custom Serve Mux
type Router struct {
	trie *types.Trie
}

// NewRouter creates a new Router for our server
func NewRouter() *Router {
	return &Router{
		trie: types.NewTrie(),
	}
}

// Get Binds the handleFunc which just handle the GET request to the pattern.
func (router *Router) Get(pattern string, handleFunc func(*Context)) {
	h, err := router.trie.Get(pattern)
	var handler HandlerInterface
	var ok bool
	if err != nil {
		handler = NewHandler()
	} else {
		handler, ok = h.(HandlerInterface)
		if !ok {
			fmt.Println("We got some shit to fix")
		}
	}
	handler.Get(handleFunc)
	router.trie.Put(pattern, handler)
}

// Post Binds the handleFunc which just handle the POST request to the pattern.
func (router *Router) Post(pattern string, handleFunc func(*Context)) {
	h, err := router.trie.Get(pattern)
	var handler HandlerInterface
	var ok bool
	if err != nil {
		handler = NewHandler()
	} else {
		handler, ok = h.(HandlerInterface)
		if !ok {
			fmt.Println("We got some shit to fix")
		}
	}
	handler.Post(handleFunc)
	router.trie.Put(pattern, handler)
}

// Put Binds the handleFunc which just handle the PUT request to the pattern.
func (router *Router) Put(pattern string, handleFunc func(*Context)) {
	h, err := router.trie.Get(pattern)
	var handler HandlerInterface
	var ok bool
	if err != nil {
		handler = NewHandler()
	} else {
		handler, ok = h.(HandlerInterface)
		if !ok {
			fmt.Println("We got some shit to fix")
		}
	}
	handler.Put(handleFunc)
	router.trie.Put(pattern, handler)
}

// Patch Binds the handleFunc which just handle the PATCH request to the pattern.
func (router *Router) Patch(pattern string, handleFunc func(*Context)) {
	h, err := router.trie.Get(pattern)
	var handler HandlerInterface
	var ok bool
	if err != nil {
		handler = NewHandler()
	} else {
		handler, ok = h.(HandlerInterface)
		if !ok {
			fmt.Println("We got some shit to fix")
		}
	}
	handler.Patch(handleFunc)
	router.trie.Put(pattern, handler)
}

// Delete Binds the handleFunc which just handle the DELETE request to the pattern.
func (router *Router) Delete(pattern string, handleFunc func(*Context)) {
	h, err := router.trie.Get(pattern)
	var handler HandlerInterface
	var ok bool
	if err != nil {
		handler = NewHandler()
	} else {
		handler, ok = h.(HandlerInterface)
		if !ok {
			fmt.Println("We got some shit to fix")
		}
	}
	handler.Delete(handleFunc)
	router.trie.Put(pattern, handler)
}

// ServeDir Serve static files.
func (router *Router) ServeDir(dirname string) {
	prefix := fmt.Sprintf("/%s", dirname)
	pattern := fmt.Sprintf("/%s/<filename:*>", dirname)
	fileserver := http.StripPrefix(prefix, http.FileServer(http.Dir(dirname)))
	router.Get(pattern, func(ctx *Context) {
		fileserver.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	})
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger(os.Stdout)
	logger.Success("Successfully accessed ServeHTTP function.")
	logger.Info("Method is: " + r.Method)
	handler, err := router.trie.Get(r.RequestURI)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	h := handler.(HandlerInterface)
	logger.Success("Successfully found Handler!")
	context := NewContext(w, r)
	switch r.Method {
	case consts.MethodGet:
		h.DoGet(context)
	case consts.MethodPost:
		h.DoPost(context)
	case consts.MethodPut:
		h.DoPut(context)
	case consts.MethodPatch:
		h.DoPatch(context)
	case consts.MethodDelete:
		h.DoDelete(context)
	default:
		logger.Error("Error: Invalid method called on router.")
		status := http.StatusMethodNotAllowed
		text := http.StatusText(status)
		context.ResponseWriter.WriteHeader(status)
		context.ResponseWriter.Write([]byte(text))
	}
}
