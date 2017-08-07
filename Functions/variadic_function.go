package main

import (
	"fmt"
)

func main() {

	response := getLargestValue(13, 44, 72, 11, 53, 12, 65)

	fmt.Println("Largest number is", response)

}

func getLargestValue(n ...int) int {

	largest := n[0]

	for _, i := range n {
		if i > largest {
			largest = i
		}
	}

	return largest
}
