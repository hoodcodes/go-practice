package main

import "fmt"

func main() {
	//untyped constants
	const make = "Toyota"
	const model = "Landcruiser"

	//multiple declarations
	const name, age = "Ralph", 25

	const (
		ssn    string  = "442741365"
		weight float64 = 162.7
	)

	//typed constants
	const greeting string = "Good Morning"
	const day int = 23

	fmt.Println(make, model, name, age, ssn, weight, greeting, day)
}
