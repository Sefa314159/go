package main

import "fmt"

var y = 42

func main(){
	fmt.Println("Hello, playground")
	fmt.Printf("%T\n", y)  	// type of y
	fmt.Printf("%b\n", y)	// y in binary
	fmt.Printf("%x\n", y)	// y in hex
	fmt.Printf("%#x\n", y)	// y in 0xhex

	s := fmt.Sprintf("%#x\t%b\t%x", y , y, y)
	fmt.Println(s)
	fmt.Printf("%v", y)
}
