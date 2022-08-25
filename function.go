package main

import "fmt"

func main() {
	sayHello("Bagus", "kurnia")
	fmt.Println(getHello("bagus"))
	firstName, _, lastName := getFullname()
	fmt.Println(firstName, lastName)
	a, b := getFullNamedReturned()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(sumAll(10, 1, 2, 3, 4, 5))
	slice := []int{99, 87, 58, 34, 22, 33}
	fmt.Println(sumAll(slice...))
	getBagus := getBagus
	fmt.Println(getBagus("bagus"))
	sayHelloWithFilter("dog", spamFilter)

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

// returning multiple values
func getFullname() (string, string, string) {
	return "Bagus", "kk", "Kurnia"
}

// function returned named value
func getFullNamedReturned() (firstName, lastName string) {
	firstName = "bagus"
	lastName = "kurnia"
	return firstName, lastName
}

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// function as value
func getBagus(name string) string {
	return name
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}

	return name
}

// Anonimous function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}
