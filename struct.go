package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indo"
	eko.Age = 12
	fmt.Println(eko)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     40,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Dakol", 33}
	fmt.Println(budi)
}
