package main

import "fmt"

func main() {

	// a is an array of only 5 values that are of type int
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// an array of two arrays, each of 3 elements that are of type int
	// can think of it like this when "parsing" the code
	// (var)                    <name>          [                        N                       {}
	// ^^^^^ optional var decl  ^^^^^^ var name ^ indicating an array of ^ number of elements of ^^ type
	// Where `type` can be another array or a concrete type
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
