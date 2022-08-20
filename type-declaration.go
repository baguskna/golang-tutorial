package main

import "fmt"

func main() {
	type NoKTP string

	var noktp NoKTP = "1234567890"
	fmt.Println(noktp)
}
