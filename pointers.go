
package main 

import "fmt"

// we are going to decleare pointers in go
// we need to know how to use pointer in our programming language 

// The type *T is a pointer a to T value 

//The & operator generates a pointer to its operand.

func main(){
    
     i,j:= 42,2701
     p := &i //pointer to i
	fmt.Println(*p) // read i through the pointer 
	*p = 21 //set i through the pointer 
   
   p = &j     // point to j
   *p = *p /37 // divide j through the pointer 
   fmt.Println(j) // see the new value of j
}
