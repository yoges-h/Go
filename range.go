
package main 

import "fmt"

var pow  = []int{1,2,4,8,16,32,64,128}

func main(){
	
	// Defining the Variable v because we have to store the Value of the Output in v
	

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n",i,v)
	}
}
