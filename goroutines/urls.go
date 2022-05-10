package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// using wait groups to synchronize on the go routines
var wg sync.WaitGroup

// using mutex to lock the shared resource like the fmt.Println
var mutx sync.Mutex

func sendRequest(url string) {
	defer wg.Done()
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	//locking the shared resource
	mutx.Lock()
	defer mutx.Unlock()
	fmt.Printf("[Status %d] on %s \n", res.StatusCode, url)
}

// go run urls.go google.com twitter.com github.com
func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run urls.go <url1> <url n>")
	}

	for _, url := range os.Args[1:] {
		url = "https://" + url
		go sendRequest(url)
		wg.Add(1)
	}

	wg.Wait()
}

//benchmark for synchronous
//8 seconds

//after implementing the go routines
//2 seconds
