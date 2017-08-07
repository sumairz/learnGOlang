package main

import (
	"fmt"
	"reflect"
)

// ":=" assignment only works inside of a function
// At package level either you can specify the variable type or just use "=" and assign a value GO will automatically assume the type by looking at the value
var (
	a = 10.00
	b = 3
)

func main() {
	fmt.Println("A", a, "and is type of ", reflect.TypeOf(a))

	c := int(a) + b

	fmt.Println("Sum ", c, "and is type of", reflect.TypeOf(c))
}
