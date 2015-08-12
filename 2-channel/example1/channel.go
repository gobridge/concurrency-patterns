/*
When the main function executes <-c, it will wait for a value to be sent.

When the boring function executes c <- value, it waits for a receiver to
be ready.

A sender and receiver must both be ready to play their part in the
communication. Otherwise we wait until they are.

Thus channels both communicate and synchronize.

Don't communicate by sharing memory, share memory by communicating.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Unbuffered Channel of strings.
	c := make(chan string)

	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		// Read From Channel - Blocking.
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}

	fmt.Println("You're boring: I'm leaving.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// Write to Channel.
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.

		// The write does not return until the read from main is complete.

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
