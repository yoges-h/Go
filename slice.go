
// An array has a fixed size. A slice, on the other hand, is a dynamically-sized 
// flexible view intop elements of an array. In practice, Slice are much more comman than array.

package main 

import "fmt"

func main(){

	primes := [6]int{2,3,5,7,11,13}

	var s []int  = primes[1:4]
	//  means it will include the first element of an array and it will exclude the last element of an array
	fmt.Println(s)
// this is my comment 
// this is my add
// hfuttygf


}
