

// Methods 

// GO does not have classes. 
// However you can define methods on types

// A method is a function with special reciver argument
// The reciver appers in  its own argement list between the func keyword and the method name.


package main 

import (

    "fmt"
   "math"
)


 type Vertex struct {

 	X,Y float64

 }

 func (v Vertex) Abs() float64 {

 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
 	// v.X means Value of X coordinate and v.Y means value of Y coordinate 
 }
func main(){
  

  v := Vertex{3,4}

  fmt.Println(v.Abs())


}