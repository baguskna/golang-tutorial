package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)

	const name = "John"

	fmt.Println(name)
}
