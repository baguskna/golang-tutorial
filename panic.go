package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("error dengan message", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(false)
	runApp(true)
}
