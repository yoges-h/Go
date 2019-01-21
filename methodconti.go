
// Method continued for go Programming language 

package main 

// A method is a function with special reciver argument.

import ( 
         "fmt"
        "math"
    )

// Here i am defining a type " MYFLOAT ""

type MyFloat float64

// Abs() is a method inside the type MyFloat

func (f MyFloat) Abs() float64 {

	if f<0 {

		return float64(-f)
	}
	return float64(f)
}

func main(){

    f := MyFloat(-math.Sqrt(5))
    fmt.Println(f.Abs())
}