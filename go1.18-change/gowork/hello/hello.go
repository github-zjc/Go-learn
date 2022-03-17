package main

import (
	"fmt"

	"golang.org/x/example"
	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

func main() {
    fmt.Println(Hello())
	fmt.Println(example.ToUpper("Hello"))
}