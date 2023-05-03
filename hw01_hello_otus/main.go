package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	s := "Hello, OTUS!"
	reversed := stringutil.Reverse(s)
	fmt.Print(reversed)
}
