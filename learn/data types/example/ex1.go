package main

import "fmt"

// function to input a number
func main() {

	fmt.Print("Enter a number: ")
	// variable float64 but not assigned value
	var input float64
	// Scanf are included of the package fmt,
	fmt.Scanf("%f", &input)
	// Declare output with the double of "input" variable
	output := input * 2

	fmt.Println(output)
}
