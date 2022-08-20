package main

import "fmt"

func main() {
	addition := 10 + 20
	fmt.Println(addition)

	subtraction := 10 - 20
	fmt.Println(subtraction)

	multiplication := 10 * 20
	fmt.Println(multiplication)

	division := 10 / 20
	fmt.Println(division)

	modulus := 10 % 20
	fmt.Println(modulus)

	increment := 10
	increment++
	fmt.Println(increment)

	decrement := 10
	decrement--
	fmt.Println(decrement)

	// augmentted assignment
	var a int = 10
	a += 10
	fmt.Println(a)
}
