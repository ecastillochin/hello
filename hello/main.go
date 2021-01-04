package main

import (
	"fmt"

	"github.com/ecastillochin/hello/greet"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(greet.English())
	fmt.Println(greet.Italian())
	fmt.Println(greet.Spanish())
	fmt.Println(quote.Hello())
}
