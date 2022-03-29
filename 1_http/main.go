package main

import (
	"log"
	"net/http"

	"github.com/KrishnaIyer/goexamples/1_http/pkg/server"
	"github.com/gorilla/mux"
)

func main() {
	address := "localhost:8080"
	r := mux.NewRouter()
	srv := server.New()

	r.HandleFunc("/", srv.HandleIndex)
	r.HandleFunc("/users/create", srv.HandleCreateUser)
	r.HandleFunc("/users/{name}", srv.HandleUser)

	s := &http.Server{
		Addr:    address,
		Handler: r,
	}

	log.Printf("Listening on port: %v", address)
	log.Fatal(s.ListenAndServe())
}
