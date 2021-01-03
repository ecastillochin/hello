package main

import (
	"fmt"

	"github.com/ecastillochin/hello/greet"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(greet.English())
	fmt.Println(greet.Italian())
	fmt.Println(greet.Spanish())
}
