package main

import (
	"log"
	"net/http"

	"github.com/brunodrugowick/go-http-server-things/pkg/server"
)

func main() {
	apiPathHandler := server.NewDefaultPathHandlerBuilder("/api").
		WithHandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, Golang!"))
		}).Build()

	srv := server.NewDefaultServerBuilder().
		SetPort(8085).
		WithPathHandler(apiPathHandler).Build()

	log.Fatal(srv.ListenAndServe())
}
