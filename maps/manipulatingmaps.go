package main

import (
	"fmt"
)

func main() {
	authorsRank := map[string]int{
		"Mak":  60,
		"John": 95,
	}
	fmt.Println(authorsRank["Mak"])

	authorsRank["Mak"] = 70
	fmt.Println(authorsRank["Mak"])

	//adding new key and value to map
	authorsRank["Calire"] = 86
	fmt.Println(authorsRank)

	delete(authorsRank, "Mak")
	fmt.Println(authorsRank)
}
