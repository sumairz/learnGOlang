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
}
