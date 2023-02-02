package main

import "fmt"

/*	Declaration of function:
	NOTE: this need a keyword "func"

	funct <nameFunction> (parameters){
		<block of code>
	}
*/

func main() {

	/*
		Declaration of variables:
		NOTE: use the keyword "var"

		var <nameVariable> <type> = <data>

		Other form:
		NOTE: the go compiler is able to infer the type based on the literal value
		<nameVariable> := <data>

		To assign value in variable do:

		<nameVariable> = <data>
	*/

	var x string = "Hello world"
	fmt.Println(x) //output: Hello world

	/*
		Declaration of constant:
		NOTE: they are created in the same way you creeat variables, but insted of using the var keyword we use the "const" keyword

		const <nameConstant> = <data>

	*/

	const y string = "Hello, word"

	/*
		Declararion of multiples variables:
		NOTE: go also has another shorthand when you need to define multiple variables.

		var(
			<nameVariable> = <data>
			<nameVariable> = <data>
			<nameVariable> = <data>
		)
	*/

}
