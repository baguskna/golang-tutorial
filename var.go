package main

import "fmt"

func main() {
	var name string

	name = "John"
	fmt.Println("Hello,", name)

	name = "Jane"
	fmt.Println("Hello,", name)

	var gender = "male"
	fmt.Println(gender)

	age := 10
	fmt.Println(age)

	age = 20
	fmt.Println(age)

	var (
		firstName = "Bagus"
		lastName  = "kurnia"
	)

	fmt.Println(firstName + lastName)
}
