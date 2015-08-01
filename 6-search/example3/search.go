/*
Example: Google Search 2.1

Given a query, return a page of search results (and some ads).
Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

Don't wait for slow servers. No locks. No condition variables. No callbacks

Run each search in their own Goroutine but only return any searches that complete in
80 Milliseconds or less
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
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
		c <- web(query)
	}()

	go func() {
		c <- image(query)
	}()

	go func() {
		c <- video(query)
	}()
	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return results
		}
	}

	return results
}
