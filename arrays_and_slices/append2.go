package main

import (
	"fmt"
)

func main() {

	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)

	for _, value := range mySlice {
		fmt.Println(value)
	}

	newSlice := []int{4, 5, 6}
	mySlice = append(mySlice, newSlice...)

	fmt.Println(mySlice)
}
