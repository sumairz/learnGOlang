package main

import (
	"fmt"
)

func main() {

	count, sum := test(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Count:", count, " Sum:", sum)
}

// When return signature is defined no need to return just assign those variables and return blank
func test(num ...int) (count int, sum int) {

	for _, n := range num {
		sum += n
	}

	count = len(num)

	return
}
