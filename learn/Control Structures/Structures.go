package main

import "fmt"

func main() {
	/*
		---if
		NOTE: Can't put brackets after newline
		NOTE: The structure can nested with other if's
		if (boolean_expression){
			<block of code>
		}

		---if else

		if(boolean_expresion){
			<block of code>
		}else{
			<statement(s) will execute if the boolean expression is false>
		}

		---if... else if... else

			if(boolean_expresion){
				<block of code>
			}else if(boolean_expresion){
				<statement(s) will execute if the boolean expression 1 is false>
				}else{
				<statement(s) will execute if none expression is true>
			}
	*/

	ask := true

	if ask {
		fmt.Println("It's true")
	}

	/*
		---Switch
		NOTE: <Expression Switch> In expression switch a case contains expressions, which is compared against the value of the expression.
		NOTE: <Type Switch> In type switch, a case contain type which is compared against the type of specially annotated switch expression.
		NOTE: In Switch default is optional

		---Expression Switch
		switch(boolean-expression or integral){
			case boolean-expression or integral type:
				statement(s);
			case boolean-expression or integral type:
				statement(s);
			default:
				statement(s);
		}
		---Type Switch

		switch x.(type){
			case type:
				statement(s);
			case type:
				statement(s);
			default:
				statement(s);
		}

	*/

	// ---EXAMPLE OF: Expression switch
	/*
		NOTE: the expression used in a switch statement must have an integral or boolean value. If the expression is not passed then the default value is true.
		NOTE: You can have any number of case statments within a switch. Each case is followed by the value to be compared to and a colon.
	*/
	var grade string = "B"
	var marks int = 90

	// Use the variable marks
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
		// Wathever of these cases.
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	// Don't have any boolean expression
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
		// This case have 2 expression
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is %s\n", grade)

	// ---EXAMPLE OF: Type Switch
	/*
		NOTE: The expression used in a switch statement must have an variable of interface{} type
		NOTE: You can have any number of case statements within a switch. Each case is followed by to be compared to and a colon.
	*/

	var x interface{}
	// var x int8

	switch i := x.(type) {

	case nil:
		fmt.Printf("type of x :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Print("x is float64")
	case func(int) float64:
		fmt.Print("x is func(int)")
	case bool, string:
		fmt.Print("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}

	/*
		---select
		NOTE: You can have any number of case statements within a select. Each case is followed by the value to be compared to and colon
		NOTE: The type for a case must be the a communication channel operation.

		select{
		case communication clause:
			statement(s);
		case communication clasue:
			statement(s);
		default:
			statement(s);
		}

		#######IN THIS THERE WILL BE NO EXAMPLE BECAUSE I DONT HAVE THE KNOWLEDGE TO UNDERSTAND SOME THINGS.##########

		---for
		NOTE: if range is avaliable, then the for loop executes for each item in the range.

		for [condition | (init; condition; increment)| Range]{
			statement(s);
		}

		example:
		for i := 1; i<5; i++ {
			sum += i
		}
		fmt.Println(sum) //Output 10 (1+2+3+4)

		---Way to make a while loop

		n :=1
		for n<5{
			n*= 2
		}
		fmt.Println(n)

		---Way to make a infinite loop

		sum := 0
		for{
			sum++
		}
		fmt.Println(sum) //never reached

		---Way to make for-each range loop

		strings := []strings{"hello", "world"}
		for i, s := range strings{
			fmt.Println(i, s)
		}

		---Way to make for-each range loop
		sum := 0
		for i := 1; i<5; i++{
			if i%2 != 0{
				continue
			}
			sum += i
		}

		fmt.Println(sum) //6 (2+4)
	*/

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}
	//
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n ", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("value of x = %d |ta %d\n", x, i)
	}
}
