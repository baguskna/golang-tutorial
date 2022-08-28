package main

import "fmt"

func returnString() interface{} {
	return 20
}

func main() {
	result := returnString()
	// resultToString := result.(string)
	// fmt.Println(resultToString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown type")
	}
}
