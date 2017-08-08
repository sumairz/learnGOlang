package main

import (
	"fmt"
)

func main() {
	authorsRank := map[string]int{
		"Mak":  60,
		"John": 95,
	}

	// Maps are unordered list it will loop randomly
	for key, value := range authorsRank {
		fmt.Printf("\nKey: %v Value is: %v", key, value)
	}

}
