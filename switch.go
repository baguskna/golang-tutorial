package main

import "fmt"

func main() {
	name := "kurnia"

	switch name {
	case "bagus":
		fmt.Println("hello bagus")
	case "kurnia":
		fmt.Println("hello kurnia")
	default:
		fmt.Println("default")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("name too long")
	case false:
		fmt.Println("name is correct")
	}
}
