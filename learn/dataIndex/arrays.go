package main

import "fmt"

func main() {
	/*
		---Arrays
		NOTE:

		var <nameVariable>[index]<dataType>
	*/

	//	NOTE: This is a way of create a array
	// var x [5]float64
	// NOTE: This is a way of assigning values to an array
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83

	// NOTE: This other way to create a array
	// x := [5]float64{98, 93, 77, 82, 83}

	// NOTE: This other way to create a array
	x := [...]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	//NOTE: compiler won't allow you create variables that you never use. Because we don't use i inside of our loop, a single underscore (_) is used to tell the compiler that we don't need this
	// for i, value := range x { WRONG FORM
	for _, value := range x {
		total += value
	}
	//Cast the variable len(x) to be the same type to total
	fmt.Println(total / float64(len(x)))

}
