package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://placeholder.it/2000*2000", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	imgData, err := ioutil.ReadAll(res.Body)

	fmt.Println("Downloaded image of size ", len(imgData))
}
