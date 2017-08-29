package main

import (
	"fmt"
	"time"
)

// learning from: https://github.com/golang/go/wiki/RateLimiting
// TODO: make this code more clearly..

func main() {
	rate := time.Second * 1
	burstLimit := 10
	tick := time.NewTicker(rate)

	// faked requests, we said the value is request id, haha..
	requests := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	}

	defer tick.Stop()

	throttle := make(chan time.Time, burstLimit)
	go func() {
		for t := range tick.C {
			select {
			case throttle <- t:
			default:
			}
		} // exits after tick.Stop()
	}()

	for _, req := range requests {
		<-throttle // rate limit our Service.Method RPCs

		go func() {
			// client.Call("Service.Method", req, ...)
			fmt.Println("faked to call service with req: ", req)
		}()
	}
}
