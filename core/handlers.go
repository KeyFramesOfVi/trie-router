package core

import (
	"fmt"
	"net/http"

	"github.com/victor-cabrera/fe-calc/server/router"
)

func indexHandler(context *router.Context) {
	// http.ServeFile(context.ResponseWriter, context.Request, filepath.Join(buildDir, "index.html"))
	context.ResponseWriter.Write([]byte("Lord Knight lives off his legacy! GET"))
}

func postHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("Lord Knight lives off his legacy! POST"))
}

func putHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("Lord Knight lives off his legacy! PUT"))
}

func patchHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("Lord Knight lives off his legacy! PATCH"))
}

func deleteHandler(context *router.Context) {
	context.ResponseWriter.Write([]byte("Lord Knight lives off his legacy! DELETE"))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		fmt.Println(r.RequestURI)
	})
}
