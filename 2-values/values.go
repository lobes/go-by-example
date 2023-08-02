package main

import "fmt"

func main() {
	// Strings can be added together with the + operator
	fmt.Println("go" + "lang")

	// Integers and floats can be added together with the + operator
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans are evaluated with the usual boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
