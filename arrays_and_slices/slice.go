package main

import (
	"fmt"
)

func main() {

	// Len is the total number of values in the slice
	// Capacity is the max number of values exist in the slice if reached max value it will double the capacity automatically

	// Slice is just referene to data in the array
	// Slice does not save values
	// Slice automatically work by reference to the main array
	// If change in main array slie will also reflect the changed value
	myCourse := make([]string, 5, 10)
	fmt.Printf("Length is %d. \nCapacity is %d \n", len(myCourse), cap(myCourse))

	myNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(myNumbers)

	// the lhs number is inclusive while RHS number is exclusive hence it will return 2,3,4
	sliceOfMyNumbers := myNumbers[2:5]
	fmt.Println(sliceOfMyNumbers)

	// This will  return values from index 0 till 4
	newSlice := myNumbers[:5]
	fmt.Println(newSlice)

	// This will  return values from index 4 to last index
	anotherSlice := myNumbers[4:]
	fmt.Println(anotherSlice)
}
