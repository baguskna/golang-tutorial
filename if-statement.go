package main

import "fmt"

func main() {
	name := "kurni"

	if name == "bagus" {
		fmt.Println(name)
	} else if name == "kurnia" {
		fmt.Println(name)
	} else {
		fmt.Println("no")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("already correct")
	}
}
