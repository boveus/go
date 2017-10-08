package main //struct - similar to an array
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	//struct fields are accesed with a dot
	v.X = 4
	fmt.Println(Vertex{1, 2})
}

//map - similar to a ruby hash
func map_example() {
	color := make(map[string]string)

	color["R"] = "Red"
	color["B"] = "Blue"
	color["G"] = "Green"

	fmt.Println(color["B"])
}

//map can behave similar to a ruby hash with key value pairs
type LatLong struct {
	Lat, Long float64
}

var m map[string]LatLong

func key_value() {
	m = make(map[string]LatLong)
	m["Bell Labs"] = LatLong{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func mutating_maps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

//array
//arrays cannot be resized in go, but you can use a slice
func array_example() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

//slices - modifying a slice modifies the underlying array
//the low bound defaults to 0 and the high pount defaults to array length
func slice_example() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	//this grabs the elements in indicies 1-4 as a slice
	var s []int = primes[1:4]
	fmt.Println(s)
}

//iterating over a slice
var power = []int{1, 2, 4, 8, 16, 32, 64, 128}

func iteration_example() {
	for i, v := range power {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
