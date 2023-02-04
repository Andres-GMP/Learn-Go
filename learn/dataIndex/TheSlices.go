package main

import "fmt"

func main() {
	fmt.Printf("Slices are great \n")

	/*
		---Slices
		NOTE: The only differencie betwen this an array is the missing length between the brackets.
		NOTE: if the lenght is zero, we can't assign value in a position
		NOTE: the slice can't have two types of variables
	*/

	//NOTE: 1st way to make slice
	// edades := []uint8{"manzana", "peras"}

	//NOTE: 2nd way to make slice
	//NOTE: We use <make> to keep memory for underlying array

	edades := make([]uint8, 0, 10)
	capacidad := cap(edades)

	fmt.Printf("Capacidad inicial: %d", capacidad)

	for i := 0; i < 25; i++ {
		edades = append(edades, i)

		if cap(edades) != capacidad {
			capacidad = cap(edades)
			fmt.Printf("La capacidad es ahora: %d", capacidad)
		}
	}
	// edades = append(edades, 22) //return a new slice assigned in slice edades with a value 22
	// edades = edades[0:6]        //resize slice
	// edades[5] = 28              //assign value index 5 of slice resize

	/*
		---Resize slice
		<sliceName> = <sliceName>[index:index]

		---Get capacity of slice
		cap(sliceName)

		---Get lenght of slice
		len(sliceName)

	*/

	fmt.Println(edades)

}
