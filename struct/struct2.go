package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Class int
}

func main() {

	sumair := Student{
		Name:  "Sumair",
		Age:   30,
		Class: 12,
	}

	fmt.Println(sumair)
}
