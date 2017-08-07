package main

import (
	"fmt"
)

func main() {
	firstRank := "614"
	secondRank := "614"

	if firstRank < secondRank {
		fmt.Println("First Rank")
	} else if firstRank > secondRank {
		fmt.Println("Second Rank")
	} else {
		fmt.Println("Same rank")
	}
}
