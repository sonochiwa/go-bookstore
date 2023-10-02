package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sonochiwa/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
