package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx_value := addValue(ctx)
	readValue(ctx_value)
}

func addValue(c context.Context) context.Context {
	return context.WithValue(c, "key name", "{test_data: Armara Wamboi}")
}

func readValue(c context.Context) {
	value := c.Value("key name")
	fmt.Println(value)
}
