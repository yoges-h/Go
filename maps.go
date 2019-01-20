package main 


import "fmt"

type Vertex struct {
	
	Lat, Long float64;

}

var m map[string]Vertex 

func main() {
	
	// The make function returns a map of the given type
	// initialized and ready for use.

	m = make(map[string]Vertex) // making map of strings of Vertex
	m["Bell Labs"] = Vertex{ // Defining the data which we are willing to map

	 40.68433,-74.399967,

	}
	fmt.Println(m["Bell Labs"]) // printing mapped data
}