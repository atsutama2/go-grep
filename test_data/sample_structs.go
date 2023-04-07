package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Car struct {
	Make  string
	Model string
	Year  int
}

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func main() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	c := Car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,
	}

	b := Book{
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		ISBN:   "1234567890",
	}

	fmt.Printf("Person: %+v\n", p)
	fmt.Printf("Car: %+v\n", c)
	fmt.Printf("Book: %+v\n", b)
}
