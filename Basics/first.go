package main

import (
	"fmt"
	"runtime"
)

var (
	name, course string
	module       float64
)

func main() {
	fmt.Println("Hello from ", runtime.GOOS)
	fmt.Println("Name ", name)
	fmt.Println("Course", course)
	fmt.Println("Module ", module)
}
