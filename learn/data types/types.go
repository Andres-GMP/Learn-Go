package main

import "fmt"

/*
	---Declaration of function:
	NOTE: this need a keyword "func"

	funct <nameFunction> (parameters){
		<block of code>
	}
*/

func main() {

	/*
		---Declaration of variables:
		NOTE: use the keyword "var"

		var <nameVariable> <type> = <data>

		---Other form:
		NOTE: the go compiler is able to infer the type based on the literal value
		<nameVariable> := <data>

		---Declaration of variables with comma:
		var <nameVariable>, <nameVariable> <type>;

		---To assign value in variable do:

		<nameVariable> = <data>

		---Declararion of multiples variables:
		NOTE: go also has another shorthand when you need to define multiple variables.

		var(
			<nameVariable> = <data>
			<nameVariable> = <data>
			<nameVariable> = <data>
		)

	*/

	var x string = "Hello world"
	fmt.Println(x) //output: Hello world

	/*
		---Types of variables:
		NOTE: <uint> means "unsigned integer" & <int> means "signed integer"
		NOTE: <uint> only contain positive numbers (or zero).
		NOTE: <byte> is the same as uint8
		NOTE: <rune> is the same as uint32
		NOTE: <uintptr> an unsigned integer to store the uninterpreted bits of a pointer value
	*/

	var eightBit uint8 = 255                       //Range: 0 to 255
	var sixteenBit uint16 = 65535                  //Range: 0 to 65535
	var thirtyTwoBit uint32 = 4294967295           //Range: 0 to 4294967295
	var sixtyFourBit uint64 = 18446744073709551615 //Range: 0 to 18446744073709551615

	fmt.Println("\n Unsigned types: ")
	fmt.Println("higer range of uint8:", eightBit)
	fmt.Println("higer range of uint16:", sixteenBit)
	fmt.Println("higer range of uint32:", thirtyTwoBit)
	fmt.Println("higer range of uint64:", sixtyFourBit)

	var signedEightBit int8 = 127                      //Range: -128 to 127
	var signedSixteenBit int16 = 32767                 //Range: -32768 to 32767
	var signedThirtyTwoBit int32 = 2147483647          //Range: -2147483648 to 2147483647
	var signedSixtyFourBit int64 = 9223372036854775807 //Range: -9223372036854775808 to 9223372036854775807

	fmt.Println("\n Signed types: ")
	fmt.Println("higer range of int8:", signedEightBit)
	fmt.Println("higer range of int16:", signedSixteenBit)
	fmt.Println("higer range of int32:", signedThirtyTwoBit)
	fmt.Println("higer range of int64:", signedSixtyFourBit)

	var floatThirtyTwoBit float32 = 1.000   //Range: IEEE-754 32-bit floating-point numbers
	var floatSixtyFourBit float64 = 1.00001 //Range: IEEE-754 64-bit floating-point numbers

	fmt.Println("\n Float types: ")
	fmt.Println("higer range of float32: are higher than this->", floatThirtyTwoBit)
	fmt.Println("higer range of float32: are higher than this->", floatSixtyFourBit)

	var complexSixtyFourBit complex64 = 123123           //Complex numbers with float32 real and imaginary parts
	var complexOneHundredTwentyEight complex128 = 123123 //Complex numbers with float64 real and imaginary parts

	fmt.Println("\n Complex types: ")
	fmt.Println("higer range of complex64: are higher than this-> ", complexSixtyFourBit)
	fmt.Println("higer range of complex128: are higher than this->", complexOneHundredTwentyEight)

	/*
		---Declaration of constant:
		NOTE: they are created in the same way you creeat variables, but insted of using the var keyword we use the "const" keyword

		const <nameConstant> = <data>
	*/

	const y string = "Hello, word"

}
