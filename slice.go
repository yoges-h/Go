
// An array has a fixed siz

package main 

import "fmt"

func main(){
// this is the just testing of vim terminal 

	primes := [6]int{2,3,5,7,11,13}

	var s []int  = primes[1:4]
	// it means it will include the first element of an array and it will exclude the last element of an array
	fmt.Println(s)

}
