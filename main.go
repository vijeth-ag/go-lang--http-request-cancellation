package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vijeth-ag/movie-ticket-booking/handlers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/movie/{name}/ticket/actions/search", handlers.MovieTicketHandler).Methods(http.MethodGet)

	http.ListenAndServe(":4000", r)
}
