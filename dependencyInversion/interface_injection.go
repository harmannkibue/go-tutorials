package main

import "fmt"

type Database interface {
	getUser() string
}

type DefaultDatabase struct{}

type FamousDatabase struct{}

func (f FamousDatabase) getUser() string {
	return "Har-Man kibue"
}

func (f FamousDatabase) printAllUsers() []string {
	return []string{
		"Harman ", "Armara", "Wamboi", "Joyce",
	}
}

func (d DefaultDatabase) getUser() string {
	return "Armara Wamboi"
}

type Greeter interface {
	sayHello(name string)
}

type NiceGreeter struct{}

func (g NiceGreeter) sayHello(name string) {
	fmt.Printf("The name passed is: %s\n", name)
}

type Program struct {
	db Database
	gr Greeter
}

func (p Program) Execute() {
	user := p.db.getUser()
	p.gr.sayHello(user)
}

func main() {
	db := DefaultDatabase{}
	db2 := FamousDatabase{}
	fmt.Println(db2.printAllUsers())
	gr := NiceGreeter{}

	p := Program{db: db, gr: gr}

	p.Execute()

}
