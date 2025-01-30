package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/brunodrugowick/go-http-server-things/pkg/server"
)

func main() {
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	log.Printf("Starting server on port %d", port)

	apiPathHandler := server.NewDefaultPathHandlerBuilder("/api").
		WithHandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, Golang!"))
		}).Build()

	srv := server.NewDefaultServerBuilder().
		SetPort(port).
		WithPathHandler(apiPathHandler).Build()

	log.Fatal(srv.ListenAndServe())
}
