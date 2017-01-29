package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type people []Person

func calculateBirthYear(age int) int {
	return 2017 - age
}

func main() {

	persons := people{
		{Name: "Agnes", Age: 19},
		{Name: "Maria", Age: 29},
	}

	for _, k := range persons {
		fmt.Println(k.Name + " you were born in " + fmt.Sprint(calculateBirthYear(k.Age)))
	}
}


