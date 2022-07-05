package main

import (
	"tutorials/UberDi"
)

func main() {
	config := UberDi.NewConfig()

	db, err := UberDi.ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	personRepository := UberDi.NewPersonRepository(db)

	personService := UberDi.NewPersonService(config, personRepository)

	server := UberDi.NewServer(config, personService)

	server.Run()
}
