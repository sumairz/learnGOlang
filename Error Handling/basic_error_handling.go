package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("c:\\temp.txt") // fake file does not exist

	if err != nil {
		fmt.Println("Error:", err)
	}
}
