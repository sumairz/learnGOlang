package main

import (
	"fmt"
)

func main() {

	Addfunc := func(num ...int) (count int, sum int) {

		for _, n := range num {
			sum += n
		}

		count = len(num)

		return
	}

	val := []int{1, 2, 3, 4, 5, 6, 7}
	count, sum := Addfunc(val...) // when sending multiple values add ... at the end
	fmt.Println("Count:", count, " Sum:", sum)
}
