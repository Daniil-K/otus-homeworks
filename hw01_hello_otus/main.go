package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

var inputString = "Hello, OTUS!"

func main() {
	fmt.Println(stringutil.Reverse(inputString))
}
