package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your host")
	user := flag.String("user", "root", "Put your user")
	password := flag.String("password", "root", "Put your password")

	flag.Parse()

	fmt.Println(*host, *user, *password)
}
