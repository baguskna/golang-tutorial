package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Loop no", counter)
		counter += 1
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("loop2 no", counter2)
	}

	slice := []string{"Bagus", "Kurnia", "Cool"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range slice {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["namme"] = "bagus"
	person["age"] = "20"

	for key, value := range person {
		fmt.Println(key, value)
	}
}
