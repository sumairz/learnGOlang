package main

import (
	"fmt"
	"os"
)

func main() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// Windows only: get loggedin username
	fmt.Println("Hello", os.Getenv("USERNAME"))
}
