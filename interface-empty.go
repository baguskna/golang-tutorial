package main

import "fmt"

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	}
	return "ups"
}

func main() {
	interger := ups(1)
	boolean := ups(2)
	strings := ups(3)

	fmt.Println(interger)
	fmt.Println(boolean)
	fmt.Println(strings)
}
