package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Printf let's us interpolate the values using some placeholder strings
	fmt.Printf("\nTypeof b: %T and c: %T\n", b, c)

	var d = true
	fmt.Println(d)

	var e int
	// interesting, it uses a default of 0 for ints
	// > Variables declared without a corresponding initialization
	// > are zero-valued. For example, the zero value for an int is 0.
	fmt.Println(e)
	fmt.Printf("Type of e: %T\n", e)

	f := "apple"
	fmt.Println(f)
}
