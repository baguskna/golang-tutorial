package main

import "fmt"

func logging() {
	fmt.Println("logging called")
}

func runApplication(value int) {
	defer logging() // will be called even the bellow return error in the end
	fmt.Println("run application")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApplication(10)
	runApplication(0)
}
