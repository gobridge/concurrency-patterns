/*
We execute a separate Goroutine to run the boring function.

The program terminates immediately because when main terminates,
the program terminates.
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

	// Program will terminate immediately.
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
