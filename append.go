
// appending a slice 
package main 

import "fmt"

func main() {

	var s []int
	printSlice(s)

	// Append the work on nill slice

	s = append(s,0)
	printSlice(s) 
    
    // The Slice grows as needed 

    s = append(s,1)
    printSlice(s)

    // We can add more than one element at a time

    s = append(s,2,3,4)
    printSlice(s)

}

func printSlice(s []int){
	fmt.Println("len = %d cap = %d %v\n",len(s), cap(s),s)
}
