package main

import "fmt"

func var_zero_values() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("These are the zero values for Go data types")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func short_var_declarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

func converting_between_datatypes() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)
}

func implicit_typing_of_vars() {
	var i int
	j := i // j is an int

	integer := 42               // int
	float := 3.142              // float64
	big_decimal := 0.867 + 0.5i // complex128

	fmt.Println(j, integer, float, big_decimal)
}

const Pi = 3.14

func declaring_constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
