/*
Restoring Sequencing

Send a channel on a channel, making Goroutine wait its turn

Receive all messages, then enable them again by sending on a private channel

Each speaker must wait for a go-ahead.

Joe: Send and Wait ---                 --- Joe: Display and Wait
                     \               /
                      ----- FanIn ---
                     /               \
Ann: Send and Wait ---                 --- Ann: Display then Release Waits
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Message contains a channel for the reply.
type Message struct {
	str  string
	wait chan bool // Acts as a signaler
}

func main() {

	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		msg1 := <-c // Waiting on someone (Joe) to talk
		fmt.Println(msg1.str)

		msg2 := <-c // Waiting on someone (Ann) to talk
		fmt.Println(msg2.str)

		msg1.wait <- true // Joe can run again
		msg2.wait <- true // Ann can run again
	}

	fmt.Println("You're boring: I'm leaving.")
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message) // The FanIn channel.

	go func() { // This Goroutine will receive messages from Joe.
		for {
			c <- <-input1 // Write the message to the FanIn channel, Blocking Call.
		}
	}()

	go func() { // This Goroutine will receive messages from Ann.
		for {
			c <- <-input2 // Write the message to the FanIn channel, Blocking Call.
		}
	}()

	return c
}

func boring(msg string) <-chan Message { // Returns receive-only (<-) channel of strings.
	c := make(chan Message)
	waitForIt := make(chan bool) // Give main control over our execution.

	go func() { // Launch the goroutine from inside the function. Function Literal.
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-waitForIt // Block until main tells us to go again.
		}
	}()

	return c // Return the channel to the caller.
}
