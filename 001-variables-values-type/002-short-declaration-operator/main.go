package main

import "fmt"

func main() {
	// ':=' declares a variable and signs a value.
	x := 42  // this is a statement.
	fmt.Println(x)
	// a statement is a small standalone element of a program that expresses some action to be carried out.
	// if i wanted to reassign a new value to x, i just use '='.
	x = 99  // because it's already been declared.
	fmt.Println(x)
	y := 100 + 24  // this is an expression. '+' is an operator, '100' and '24' are the operands.
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
}
