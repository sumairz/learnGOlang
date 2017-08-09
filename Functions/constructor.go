package main

import (
	"fmt"
)

func main() {
	foo := myStructConstructor()
	foo.myMap["name"] = "Sumair"

	fmt.Println(foo)
}

type myStruct struct {
	myMap map[string]string
}

func myStructConstructor() *myStruct {
	result := myStruct{}
	result.myMap = map[string]string{}

	return &result
}
