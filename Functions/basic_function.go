package main

import (
	"fmt"
	"strings"
)

func main() {

	welcome := "hello, sumair"
	bye := "bye, sumair"

	fmt.Println(toTitleCase(welcome, bye))
}

func toTitleCase(text1, text2 string) (r1, r2 string) {

	text1 = strings.ToUpper(text1)
	text2 = strings.ToUpper(text2)

	return text1, text2
}
