package main

import (
	"fmt"
)

func main() {

	myArray := [3]int{}
	myArray[0] = 1
	myArray[1] = 22
	myArray[2] = 45

	fmt.Println(myArray)

	newArr := [3]int{1, 244, 567}
	fmt.Println(newArr)

	newArr[2] = 11
	fmt.Println(newArr)

	strArr := [5]string{"Sumair", "Haris", "Askani"}
	strArr[3] = "Alex"
	fmt.Println(strArr)

	// array size based on the length of data
	lenArr := [...]int{1, 2, 4, 5, 6, 7, 8}
	//lenArr[7] = 100 // in this case cannot add new index to the array
	fmt.Println(lenArr)

	// Creating a slice of an array
	mySlice := lenArr[2:3]
	fmt.Println("Slice: ", mySlice, " and length is:", len(mySlice))
}
