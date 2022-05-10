package main

import (
	"fmt"
	"strconv"
)

type Human interface {
	getInfo() string
	setInfo()
}
type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name    string
	Subject string
}

func (s Student) getInfo() string {
	ret := s.Name + " of age " + strconv.Itoa(s.Age)
	return ret
}

func (s *Student) setInfo() {
	fmt.Println("Enter students information: ")
	fmt.Scan(&s.Name)
	fmt.Scanln(&s.Age)
}

func (t *Teacher) setInfo() {
	fmt.Println("Enter teachers information: ")
	fmt.Scanln(&t.Name)
	fmt.Scanln(&t.Subject)
}

func (t *Teacher) getInfo() string {
	ret := t.Name + " teaches subject called " + t.Subject
	return ret
}

func main() {
	human := []Human{new(Student), new(Teacher)}

	for _, hum := range human {
		hum.setInfo()
		fmt.Println(hum.getInfo())
	}

}
