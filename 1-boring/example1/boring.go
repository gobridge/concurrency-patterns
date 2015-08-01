// Program loops five times printing a message.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boring("boring!")
}

func boring(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
