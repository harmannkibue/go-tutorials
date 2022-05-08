package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//creating a context
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*10000)
	defer cancel()

	//	create a http request
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://via.placeholder.com/150", nil)
	if err != nil {
		panic(err)
	}

	// perform http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	image_data, err := ioutil.ReadAll(res.Body)
	fmt.Printf("Downloaded image of size %d ", len(image_data))

}
