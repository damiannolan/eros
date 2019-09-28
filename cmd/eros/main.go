package main

import (
	"fmt"

	"github.com/damiannolan/eros/server"
)

func main() {
	fmt.Println("Eros")

	srv := &server.Server{
		Name: "Http Server Struct",
	}

	fmt.Println(srv)
}
