/*
Example: Google Search 2.1 - Avoid Timeouts

Given a query, return a page of search results (and some ads).
Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

No locks. No condition variables. No callbacks

Replicate the servers. Send requests to multiple replicas, and use the first response.

Run the same search against multiple servers in their own Goroutine
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	result string
	search func(query string) result
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	// Run the search against two replicas
	result := first("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))

	elasped := time.Since(start)

	fmt.Println(result)
	fmt.Println(elasped)
}

func fakeSearch(kind string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func first(query string, replicas ...search) result {
	c := make(chan result)

	// Define a function that takes the index to the replica function to use.
	// Then it executes that function writing the results to the channel.
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	// Run each replica function in its own Goroutine.
	for i := range replicas {
		go searchReplica(i)
	}

	// As soon as one of the replica functions write a result, return.
	return <-c
}
