package main //struct - similar to an array
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
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
