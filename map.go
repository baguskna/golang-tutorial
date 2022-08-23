package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "John",
		"age":  "30",
	}

	person["title"] = "software engineer"

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "The Gopher"
	book["author"] = "John Doe"
	book["publisher"] = "O'Reilly"

	fmt.Println(book)

	delete(book, "publisher")

	fmt.Println(book)
}
