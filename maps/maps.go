package main

import (
	"fmt"
)

func main() {

	// Maps passed to functions by reference
	// Changes in original maps will reflect everywhere it is used
	// Unsfe for concurrency

	// Map Format: make(map[KeyType]ValueType, size)
	// specifying size can improve performance when working with big maps
	articleRank := make(map[string]int)
	articleRank["article1"] = 67
	articleRank["article2"] = 80

	authorsRank := map[string]int{
		"Mak":  60,
		"John": 95,
	}

	fmt.Printf("Articles Rank: %v\n Authors Rank: %v", articleRank, authorsRank)

}
