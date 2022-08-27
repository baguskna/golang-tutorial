package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	eko := Person{
		Name: "Eko koko",
	}

	SayHello(eko)
}
