/*
Example: Google Search

Given a query, return a page of search results (and some ads).
Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

Google function takes a query and returns a slice of Results (which are just strings)
Google invokes Web, Image and Video searches serially, appending them to the results slice.

Run each search in series
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
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))

	return results
}
