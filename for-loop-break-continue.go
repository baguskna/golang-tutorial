package main

import "fmt"

func main() {
	// break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("loop", i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("loop not continue", i)
	}
}
