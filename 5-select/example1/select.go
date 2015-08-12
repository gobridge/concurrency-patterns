/*
Select is a control structure that is unique to concurrency.

The reason channels and Goroutines are built into the language.

Like a switch but each case is a communication:
-- All channels are evaluated
-- Selection blocks until one communication can proceed, which then does.
-- If multiple can proceed, select choose pseudo-randomly.
-- Default clause, if present, executes immediately if no channel is ready.

Multiplexing: Let whosoever is ready to talk, talk.

The fanIn function fronts the other channels. Goroutines that are ready to talk
can independently talk without Blocking the other Goroutines. The FanIn channel
receives all messages for processing.

Decouples the execution between the different Goroutines.

Joe ---
       \
        ----- FanIn --- Independent Messages Displayed
       /
Ann ---
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // Display any message received on the FanIn channel.
	}

	fmt.Println("You're boring: I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string) // The FanIn channel

	go func() { // Now using a select and only one Goroutine
		for {
			select {
			case s := <-input1:
				c <- s

			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
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
