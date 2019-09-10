// packages is a way for us to organize our code
package main

import "fmt"

func main() {
	// from package fmt, i am calling print line
	fmt.Println("Hello World !", 42, true)

	foo()

	fmt.Println("something more")

	for i := 0; i < 10; i++{
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo(){
	fmt.Println("I'm in foo")
}

func bar(){
	fmt.Println("I'm in bar")
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
