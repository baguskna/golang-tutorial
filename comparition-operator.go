package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	if a < b {
		fmt.Printf("a is less than b\n")
	} else {
		fmt.Printf("a is not less than b\n")
	}

	name1 := "John"
	name2 := "John"

	fmt.Println(name1 == name2)

	name3 := "Eko"
	name4 := "Budi"

	fmt.Println(name3 > name4) // will return 'true' because "Eko" is alphabetically after "Budi"
}
