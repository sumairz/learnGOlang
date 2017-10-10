package main

import (
	"fmt"
)

type Class []Student

func (c *Class) AddStudent(st Student) {
	*c = append(*c, st)
}

func (c *Class) Count() int {
	return len(*c)
}

type Student struct {
	Name  string
	Age   int
	Class int
}

func (this *Student) SetParam(param string, value interface{}) {

	switch param {
	case "name":
		this.Name = value.(string)
	case "age":
		this.Age = value.(int)
	case "class":
		this.Class = value.(int)
	}
}

func (this *Student) GetParam(param string) interface{} {
	switch param {
	case "name":
		return this.Name
	case "age":
		return this.Age
	case "class":
		return this.Class
	default:
		return nil
	}
}

func main() {
	var (
		s1, s2 Student
		cl     Class
	)

	s1.SetParam("name", "Sumair")
	s1.SetParam("age", 15)
	s1.SetParam("class", 10)

	s2.SetParam("name", "Haris")
	s2.SetParam("age", 14)
	s2.SetParam("class", 9)

	cl.AddStudent(s1)
	cl.AddStudent(s2)

	fmt.Println(cl.Count())
}
