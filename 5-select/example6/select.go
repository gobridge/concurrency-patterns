/*
Daisy-Chain

1 ->    2          3          4          5          6          7
     Gopher1 -> Gopher2 -> Gopher3 -> Gopher4 -> Gopher5 -> Gopher6
*/
package main

import (
	"fmt"
)

func main() {
	const n = 6
	rightmost := make(chan int)

	var left chan int
	right := rightmost

	// Launch Goroutines that each act as a gopher waiting for a message.
	for i := 0; i < n; i++ {
		left = make(chan int)
		go pass(left, right) // Gopher6 passes to Gopher7, Gopher5 passes to Gopher6, etc.
		right = left
	}

	fmt.Println("Goroutines Are Waiting")

	// Send the first message to Gopher1.
	go func(c chan int) {
		fmt.Println("Give Gopher1 the inital value")
		c <- 1
	}(left)

	// Wait for the message to reach the end.
	fmt.Printf("Final Value: %d\n", <-rightmost)
}

func pass(left, right chan int) {
	value := <-left
	right <- 1 + value

	fmt.Printf("Left[%d] Right[%d]\n", value, value+1)
}
