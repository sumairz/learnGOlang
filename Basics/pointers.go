package main

import (
	"fmt"
)

var (
	a = 3.2
)

func main() {

	ptr := &a

	fmt.Println("Memory location of a ", &a)
	fmt.Println("Value of a ", *ptr)
}
