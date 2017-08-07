package main

import (
	"fmt"
)

func main() {

	course := "GIT Basics"

	fmt.Println("Course name is", course)

	changeCourse(course)

	fmt.Println("New couse name is", course)
}

func changeCourse(c string) string {
	c = "Docker: Containerize your app"

	fmt.Println("Inside func course name", c)

	return c
}
