/*
We execute a separate Goroutine to run the boring function.

The program terminates immediately because when main terminates,
the program terminates.

We give the Goroutine two seconds to perform some work.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Use a goroutine.
	go boring("boring!")

	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring: I'm leaving.")

	// Program will terminate immediately.
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
