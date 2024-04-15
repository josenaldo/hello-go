package main

import "fmt"

// you can declare variables without assigning a value to them.
// In this case, the variables will have a zero value.
// The zero value of a variable is the default value assigned to it when it is declared.
// The zero value of a variable depends on its type.
// For example, the zero value of an integer is 0, and the zero value of a string is an empty string.

func main() {
	var a int
	var b string
	var c float32
	var d bool

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
}
