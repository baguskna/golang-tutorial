package main

import "fmt"

func main() {
	var nilai32 int32 = 100_000
	var nilain64 int64 = int64(nilai32)

	fmt.Println(nilai32, nilain64)

	name := "John"
	firstLetter := name[0]
	toString := string(firstLetter)
	fmt.Println(name, toString)
}
