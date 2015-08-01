/*
Quit Channel

You can turn this around and tell Joe to stop when we're tired of listening to him.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- true
	fmt.Println("EXIT")
}

func boring(msg string, quit chan bool) <-chan string { // Returns receive-only (<-) channel of strings.
	c := make(chan string)

	go func() { // Launch the goroutine from inside the function. Function Literal.
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// Do Nothing
			case <-quit:
				fmt.Println("Quiting")
				return
			}
		}
	}()

	return c // Return the channel to the caller.
}
