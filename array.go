package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// fmt.Println(names[3]) // will return 'panic: runtime error: index out of range'

	var values = [3]int{1, 2, 3}
	fmt.Println(values)
	fmt.Println(len(values))
}
