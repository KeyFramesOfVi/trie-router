package core

import (
	"fmt"
	"net/http"

	"github.com/victor-cabrera/fe-calc/server/router"
)

func indexHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("GET handler call."))
}

func postHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("POST handler call"))
}

func putHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("PUT handler call."))
}

func patchHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("PATCH handler call."))
}

func deleteHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("DELETE handler call."))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		fmt.Println(r.RequestURI)
	})
}
