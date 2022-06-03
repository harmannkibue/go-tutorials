package main

import "fmt"

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserting lighting connector into computer")
	com.insertIntoLightningPort()
}
