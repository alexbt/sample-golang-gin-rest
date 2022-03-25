package main

import (
	"fmt"
	server "github.com/alexbt/sample-golang/pkg"
)

func main() {
	fmt.Print("hello")
	server.Start()
}
