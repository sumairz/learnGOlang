package main

import (
	"fmt"
)

func main() {

	type courseMata struct {
		Author string
		Level  string
		Rating float64
	}

	//method to declare instance of struct
	//var biology courseMata // instance with zero value
	//maths := new(courseMata) // instance with zero value
	learnGOlang := courseMata{
		Author: "Sumair",
		Level:  "Beginner",
		Rating: 9,
	}

	fmt.Println("Author of this course is:", learnGOlang.Author)
	fmt.Println(learnGOlang)

}
