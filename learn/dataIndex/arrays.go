package main

import "fmt"

func main() {
	/*
		---Arrays
		NOTE:

		var <nameVariable>[index]<dataType>
	*/

	//This is a way of create a array
	var x [5]float64
	//This is a way of assigning values to an array
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	//loop to compute a the total score.
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	//
	fmt.Println(total / float64(len(x)))

}
