package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// pointer function
func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	// '&' is to pointer to address1 with the same data on memory
	// address2 has type *Address
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	// '*' swapping its reference from address1 to address3
	*address3 = Address{
		City:     "Jakarta",
		Province: "Jakarta",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	fmt.Println(address4)

	address5 := Address{
		City:     "Stockholm",
		Province: "Stockholm",
		Country:  "Sweden",
	}

	changeCountryToIndonesia(&address5)

	fmt.Println(address5)
}
