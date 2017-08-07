package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "Sumair"
	course = "OOP"
	module = 3.2
)

func main() {
	fmt.Println("Name ", name, "and is type of ", reflect.TypeOf(name))
	fmt.Println("Course", course, "and is type of ", reflect.TypeOf(course))
	fmt.Println("Module ", module, "and is type of", reflect.TypeOf(module))
}
