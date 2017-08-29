package main

import (
	"fmt"
)

func generate(ch chan<- int, n int) {
	for i := 2; i < n; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}

	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}

	close(dst)
}

func sieve() {
	var n = 100
	ch := make(chan int) // Create a new channel.
	go generate(ch, n)

	for {
		prime, ok := <-ch
		if !ok {
			fmt.Println("Run over.")
			break
		}

		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
