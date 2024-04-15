package main

import "fmt"

// Variables are used to store data in memory.
// They can be of different types, such as integers, strings, and floats.
// Here we declare some variables with different types.
// The type of a variable is defined after the variable name.
// The value of a variable is assigned using the = operator.
// This kind of declaration, with the keyword car and a pair of parentheses, is called a block declaration.
var (
	IsBiggerThan18 bool    = true
	Age            int     = 20
	Name           string  = "John"
	Height         float32 = 1.75
	Weight         float32 = 70.5
)

// You can also declare variables one by one.
var LastName string = "Doe"

// The main function is the entry point of the program.
func main() {
	// We print the values of the variables to the console.
	fmt.Println("IsBiggerThan18:", IsBiggerThan18)
	fmt.Println("Age:", Age)
	fmt.Println("Name:", Name)
	fmt.Println("LastName:", LastName)
	fmt.Println("Height:", Height)
	fmt.Println("Weight:", Weight)
}
