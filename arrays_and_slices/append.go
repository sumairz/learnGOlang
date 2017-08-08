package main

import (
	"fmt"
)

func main() {

	// append will create a new array and copy old values to the new one
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is %d. \nCapacity is %d \n", len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}
}
