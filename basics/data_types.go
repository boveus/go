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
