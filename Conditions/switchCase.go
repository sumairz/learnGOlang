package main

import (
	"fmt"
)

func main() {

	num := 2

	switch num {
	case 1:
		fmt.Println("Its number One")
	case 2:
		fmt.Println("getting next case")
		fallthrough //This will match 2 and immediately run the next case as well
	case 3:
		fmt.Println("Its number two or three")
	case 4:
	case 5:
		fmt.Println("Its number four or five")
	default:
		fmt.Println("Its number", num)
	}
}
