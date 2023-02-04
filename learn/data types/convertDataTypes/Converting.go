package main

// NOTE: Can import multiple package with parentesis
// NOTE: Can't import package and not used
import (
	"fmt"
	"strconv"
)

func main() {

	edad := "22" //variable declarated string

	//NOTE: a function can return TWO results, in these case are Atoi to return a string and int
	//NOTE: To get the 2nd return in a variable use the _
	//NOTE:
	edad_int, _ := strconv.Atoi((edad))

	fmt.Println(edad_int + 10)
}

// REFERENCES: https://youtu.be/CIIbEzXoPJI?list=PLdKnuzc4h6gFmPLeous4S0xn0j9Ik2s3Y
