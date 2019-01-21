
package main 

import ( 

"fmt"

"math"

)

// Here Vertex is Type 

type Vertex struct {

	X,Y float64

}

// Here ABs is a Funcion On Type Vertex 

func Abs(v Vertex) float64 {

	return math.Sqrt(v.X*v.X + v.Y * v.Y)

}

// Here Scale is a Function on Type Vertex

func Scale(v *Vertex, f float64) {

	v.X = v.X * f
	v.Y = v.Y * f

}

func main() {

	v := Vertex{3, 4}

	Scale(&v, 10)

	fmt.Println(Abs(v))

}