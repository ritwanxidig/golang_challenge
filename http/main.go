package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", api.getIndexHandler)
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.postUsersHandler)

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}
