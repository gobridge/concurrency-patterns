/*
Example: Google Search 3.0

Given a query, return a page of search results (and some ads).
Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

No locks. No condition variables. No callbacks

Reduce tail latency using replicated search servers

Run the same search against multiple servers in their own Goroutine but only return searches
that complete in 80 Milliseconds or less

All three searches SHOULD always come back in under 80 milliseconds
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web1   = fakeSearch("web")
	web2   = fakeSearch("web")
	image1 = fakeSearch("image")
	image2 = fakeSearch("image")
	video1 = fakeSearch("video")
	video2 = fakeSearch("video")
)

type (
	result string
	search func(query string) result
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	results := google("golang")
	elasped := time.Since(start)

	fmt.Println(results)
	fmt.Println(elasped)
}

func fakeSearch(kind string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func google(query string) (results []result) {
	c := make(chan result)

	go func() {
		c <- first(query, web1, web2)
	}()

	go func() {
		c <- first(query, image1, image2)
	}()

	go func() {
		c <- first(query, video1, video2)
	}()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return results
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
