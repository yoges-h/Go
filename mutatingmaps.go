
package main 

import "fmt"

func main(){

	m := make(map[string]int)

        // Insert or Update an element in map m
       // Retrieve an element 
      // elem = m[Key]

	m["Answer"] = 42
	fmt.Println(" The Value :",m["Answer"])

	m["Answer"] = 48
      	fmt.Println("The value :",m["Answer"])
       // delete an element 

	delete(m,"Answer")
	fmt.Println("The Value :",m["Answer"])
     // Test that a key is present with a two -value assignement:
	v, ok := m["Answer"]
	fmt.Println("The Value :",v,"Present?",ok)
// IF KEY IS PERESENT IN M, OK IT IS TRUE. IF NOT OK IS FALSE.
// IF KEY IS NOT IN MAP THEN ELEM IS THE ZERO VALUE FOR TEH MAPS ELEMENTYPE.
}
