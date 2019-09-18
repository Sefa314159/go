package main

import "fmt"

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of type int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main()  {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE(of a certain TYPE)
	x := 42  // this must be in the main, otherwise we would be getting an error.
	fmt.Println(x)
	var y = 43  // we can declare a variable outside of a function body, we use 'var'
	fmt.Println(y)

	foo()
}

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initilization
var y = 44  // the scope of y would be package level

func foo(){
	fmt.Println("var:", y)  // y = 44
	fmt.Println(z)		// z = 0
}
