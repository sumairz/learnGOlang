package main

import (
	"fmt"
)

func main() {

	names := []string{"Sumair", "Haris", "Waleed", "Umair"}

	// When we need both key and value
	for key, value := range names {
		fmt.Println(key, " --- ", value)
	}

	// When we don't need key just use "_" to ignore that value
	for _, value := range names {
		fmt.Println(value)
	}
}
