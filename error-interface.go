package main

import (
	"errors"
	"fmt"
)

func main() {
	var contohError error = errors.New("Ups error")
	fmt.Println(contohError.Error())
	hasil, err := pembagi(1, 1)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}
