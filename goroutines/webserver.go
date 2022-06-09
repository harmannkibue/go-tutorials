package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var nextID = make(chan int)
	go counter(nextID)
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Harman youve got %d</h1>", nextID)
	fmt.Println("THe id iss ", nextID)
}
