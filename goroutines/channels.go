package main

import (
	"log"
	"net/http"
	"time"
)

// This is a struct for the results
type result struct {
	url     string
	err     error
	latency time.Duration
}

// Getting data from a given url
func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	result := make(chan result)
	urlList := []string{
		"http://localhost:8080",
		"http://localhost:8080",
		"http://localhost:8080",
		"http://localhost:8080",
		"http://localhost:8080",
		"http://localhost:8080",
		//"https://google.com",
		//"https://amazon.com",
		//"https://financialhubfx.com",
		//"https://twitter.com",
	}

	for _, url := range urlList {
		go get(url, result)
	}

	for range urlList {
		r := <-result

		if r.err != nil {
			log.Printf("%-20s %s", r.url, r.err)
		} else {
			log.Printf("%-20s %s", r.url, r.latency)
		}
	}

}
