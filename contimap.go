
package main 

import "fmt"

// Map literals continued 

type Vertex  struct {
	Lat, Long float64

}

// If the top-level type is just name, you can omit from the elements of the literals.

var m = map[string]Vertex{

	"Bell Labs" : {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},

}
func main() {

 fmt.Println(m)

}