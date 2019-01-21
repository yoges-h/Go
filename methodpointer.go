package main 

import ( 
    
    "fmt"
    "math"

)

// creating a type Vertex 

type Vertex struct {

	X, Y float64 
}

// Here Vertex is Type
func (v Vertex ) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

// Here Scale is a Method on Type Vertex

func (v *Vertex) Scale(f float64) {

	v.X = v.X*f
	v.Y = v.Y*f

}

func main() {

	v := Vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())

}