package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

var inputString = "Hello, OTUS!"

func main() {
	fmt.Println(reverse.String(inputString))
}
