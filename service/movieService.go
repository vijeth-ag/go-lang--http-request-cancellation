package service

import (
	"context"
	"fmt"
	"time"

	"github.com/vijeth-ag/movie-ticket-booking/cutomtypes"
)

func moviesSearchQuery(ctx context.Context, readCh chan cutomtypes.Result) {
	time.Sleep(5 * time.Second)
	payload := cutomtypes.Result{
		Available:      true,
		SeatsAvailable: 10,
	}
	time.Sleep(5 * time.Second)
	readCh <- payload
	fmt.Println("sending result from goroutine")
}

func SearchMovies(ctx context.Context, ch chan cutomtypes.Result) {

	readCh := make(chan cutomtypes.Result)
	go moviesSearchQuery(ctx, readCh)
	for {
		select {
		case <-ctx.Done():
			// Received Done signal from parent
			fmt.Println("at Done() Stopping work")
			return
		case receivedResult := <-readCh:
			fmt.Println("result available", receivedResult)
			ch <- receivedResult
			return
		}
	}

}
