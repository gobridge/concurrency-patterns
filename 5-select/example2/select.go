/*
Timeout Using Select

The time.After function returns a channel that blocks for the specified duration.
After the interval, the channel delivers the current time, once.

The select is giving the boring routine 800ms to respond. This will be an endless
loop if boring can perform its work under 800ms every time.

*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Joe")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(800 * time.Millisecond): // This is reset on every iteration.
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
