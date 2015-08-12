/*
Timeout Using Select

Create the timer once, outside the loop, to time out the entire conversation.
(In the previous program, we had a timeout for each message)

This time the program will terminate after 5 seconds
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second) // Terminate program after 5 seconds.

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You're too slow.")
			return
		}
	}
}

func boring(msg string) <-chan string { // Returns receive-only (<-) channel of strings.
	c := make(chan string)

	go func() { // Launch the goroutine from inside the function. Function Literal.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // Return the channel to the caller.
}
