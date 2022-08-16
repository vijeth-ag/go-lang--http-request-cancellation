package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/vijeth-ag/movie-ticket-booking/cutomtypes"
	"github.com/vijeth-ag/movie-ticket-booking/service"
)

type requestContextMap struct {
	IPAddress context.Context
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func MovieTicketHandler(w http.ResponseWriter, r *http.Request) {

	enableCors(&w) //testing purpose

	ch := make(chan cutomtypes.Result)

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	go service.SearchMovies(ctx, ch)
	fmt.Println("started request process.............")

	select {
	case <-ctx.Done():
		fmt.Println("request Done signal received ")
		fmt.Fprint(os.Stderr, "request cancelled\n")
	case obj := <-ch:
		fmt.Println("obj", obj)
		json.NewEncoder(w).Encode(obj)
	}
}
